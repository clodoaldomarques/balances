package dynamodb

import (
	"balances/internal/app/domain/accounts"
	"balances/internal/commons"
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/shopspring/decimal"
)

type Account struct {
	AccountID int64              `dynamodbav:"account_id"`
	OrgID     string             `dynamodbav:"org_id"`
	Limits    commons.DecimalMap `dynamodbav:"limits"`
	Balances  commons.DecimalMap `dynamodbav:"balances"`
	CreatedAt time.Time          `dynamodbav:"created_at"`
	UpdatedAt time.Time          `dynamodbav:"updated_at"`
	Status    string             `dynamodbav:"status"`
	Version   int64              `dynamodbav:"version"`
}

type Entry struct {
	TrackingID string    `dynamodbav:"tracking_id"`
	AccountID  int64     `dynamodbav:"account_id"`
	OrgID      string    `dynamodbav:"org_id"`
	Impacts    []byte    `dynamodbav:"impacts"`
	CreatedAt  time.Time `dynamodbav:"created_at"`
}

type Impact struct {
	Balance   string          `dynamodbav:"balance"`
	Operation string          `dynamodbav:"operation"`
	Amount    decimal.Decimal `dynamodbav:"amount"`
	Rules     []string        `dynamodbav:"rules"`
}

func (a Account) GetKey() map[string]types.AttributeValue {
	acc, err := attributevalue.Marshal(a.AccountID)
	if err != nil {
		panic(err)
	}
	org, err := attributevalue.Marshal(a.OrgID)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"account_id": acc, "org_id": org}
}

func (a Account) toEntity() accounts.Account {
	return accounts.Account{
		AccountID: a.AccountID,
		OrgID:     a.OrgID,
		Limits:    a.Limits,
		Balances:  a.Balances,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		Status:    accounts.Status(a.Status),
		Version:   a.Version,
	}
}

func toAccount(a accounts.Account) (map[string]*dynamodb.AttributeValue, error) {
	acc := Account{
		AccountID: a.AccountID,
		OrgID:     a.OrgID,
		Limits:    a.Limits,
		Balances:  a.Balances,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		Status:    string(a.Status),
		Version:   a.Version,
	}
	return dynamodbattribute.MarshalMap(acc)

}

func toEntries(e accounts.Entry) (map[string]*dynamodb.AttributeValue, error) {
	imps, err := json.Marshal(toImpacts(e.Impacts))
	if err != nil {
		return nil, err
	}
	en := Entry{
		TrackingID: e.TrackingID,
		AccountID:  e.AccountID,
		OrgID:      e.OrgID,
		Impacts:    imps,
		CreatedAt:  e.CreatedAt,
	}

	return dynamodbattribute.MarshalMap(en)
}

func toImpacts(impact []accounts.Impact) []Impact {
	impacts := make([]Impact, 0, len(impact))
	for _, i := range impact {
		new := Impact{
			Balance:   i.Balance,
			Operation: i.Operation,
			Amount:    i.Amount,
			Rules:     i.Rules,
		}
		impacts = append(impacts, new)
	}
	return impacts
}
