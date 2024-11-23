package commons

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/shopspring/decimal"
)

func init() {
	decimal.MarshalJSONWithoutQuotes = true
}

type DecimalMap map[string]decimal.Decimal

func (d DecimalMap) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d *DecimalMap) Scan(src interface{}) error {
	switch data := src.(type) {
	case []byte:
		return json.Unmarshal(data, d)
	case string:
		return json.Unmarshal([]byte(data), d)
	default:
		return fmt.Errorf("unsupported type: %T", src)
	}
}

func (d *DecimalMap) MarshalJSON() ([]byte, error) {
	var m map[string][]byte
	for k, v := range *d {
		value, err := v.MarshalJSON()
		if err != nil {
			return nil, err
		}
		m[k] = value
	}
	return json.Marshal(m)
}

func (d *DecimalMap) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	fmt.Println("chamou o metodo")
	return nil, nil
}
