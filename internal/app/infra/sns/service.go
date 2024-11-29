package sns

import (
	"balances/configs"
	"balances/internal/app/domain/events"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/service/sns"
)

type Publisher struct {
	ctx      context.Context
	svc      *sns.SNS
	topicARN string
}

func NewPublisher(ctx context.Context) *Publisher {
	svc, err := Connect()
	if err != nil {
		panic(err)
	}
	return &Publisher{
		ctx:      ctx,
		svc:      svc,
		topicARN: configs.New().SnsAccountTopic,
	}
}

func (p Publisher) Emit(ctx context.Context, evt events.Event) {
	fmt.Println(*evt.ToMessage())
	result, err := p.svc.Publish(&sns.PublishInput{
		Message:  evt.ToMessage(),
		TopicArn: &p.topicARN,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*result.MessageId)
}
