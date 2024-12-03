package dynamodb

// import (
// 	"balances/internal/app/domain/accounts"
// 	"context"
// 	"strconv"
// 	"time"

// 	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
// )

// const (
// 	AccountTableName = "accounts"
// 	EntriesTableName = "entries"
// )

// type Repository struct {
// 	ctx context.Context
// 	db  *dynamodbstreams.Client
// }

// func NewRepository(ctx context.Context) *Repository {
// 	db := Connect(ctx)
// 	return &Repository{
// 		ctx: ctx,
// 		db:  db,
// 	}
// }

// func (r Repository) Close() {
// }

// func (r Repository) SaveNewAccount(ctx context.Context, a accounts.Account) error {
// 	item, err := toAccount(a)
// 	if err != nil {
// 		return err
// 	}

// 	input := &dynamodb.PutItemInput{
// 		Item:      item,
// 		TableName: aws.String(AccountTableName),
// 	}

// 	_, err = r.db.PutItemWithContext(ctx, input)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func (r Repository) UpdateExistingAccount(ctx context.Context, a accounts.Account) error {
// 	limits, err := dynamodbattribute.Marshal(a.Limits)
// 	if err != nil {
// 		return err
// 	}

// 	balances, err := dynamodbattribute.Marshal(a.Balances)
// 	if err != nil {
// 		return err
// 	}

// 	input := &dynamodb.UpdateItemInput{
// 		TableName: aws.String(AccountTableName),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"account_id": {
// 				N: aws.String(strconv.Itoa(int(a.AccountID))),
// 			},
// 			"org_id": {
// 				S: aws.String(a.OrgID),
// 			},
// 		},
// 		UpdateExpression: aws.String("set limits = :l, balances = :b, updated_at = :u, status = :s, version = :v"),
// 		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
// 			":l": limits,
// 			":b": balances,
// 			":u": {S: aws.String(a.UpdatedAt.Format(time.RFC3339Nano))},
// 			":s": {S: aws.String(string(a.Status))},
// 			":v": {N: aws.String(strconv.Itoa(int(a.Version)))},
// 		},
// 		ReturnValues: aws.String("UPDATED_NEW"),
// 	}

// 	_, err = r.db.UpdateItemWithContext(ctx, input)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func (r Repository) RetrieveAccountByID(ctx context.Context, accountID int64, orgID string) (accounts.Account, error) {
// 	item := &dynamodb.GetItemInput{
// 		TableName:      aws.String(AccountTableName),
// 		ConsistentRead: aws.Bool(true),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"account_id": {
// 				N: aws.String(strconv.Itoa(int(accountID))),
// 			},
// 			"org_id": {
// 				S: aws.String(orgID),
// 			},
// 		},
// 	}

// 	result, err := r.db.GetItemWithContext(ctx, item)
// 	if err != nil {
// 		return accounts.Account{}, err
// 	}
// 	var acc accounts.Account
// 	err = dynamodbattribute.UnmarshalMap(result.Item, &acc)
// 	if err != nil {
// 		return accounts.Account{}, err
// 	}
// 	return acc, nil
// }

// func (r Repository) SaveEntryAndUpdateAccount(ctx context.Context, e accounts.Entry, a accounts.Account) error {
// 	return nil
// }
