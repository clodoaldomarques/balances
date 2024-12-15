package configs

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	AppPort          int
	MySqlDBUser      string
	MySqlDBPass      string
	MySqlDBHost      string
	MySqlDBPort      string
	MysqlDBName      string
	AwsAddress       string
	AwsRegion        string
	AccessKeyID      string
	SecretAccessKey  string
	BalancesSNSTopic string
	BalancesSQSQueue string
}

type Option func(*Config)

func New(options ...Option) *Config {
	c := &Config{
		AppPort:          GetInt("APP_PORT", 8080),
		MySqlDBUser:      GetString("MYSQL_USER", "admin"),
		MySqlDBPass:      GetString("MYSQL_PASS", "b4l4nc3s"),
		MySqlDBHost:      GetString("MYSQL_HOST", "192.168.49.2"),
		MySqlDBPort:      GetString("MYSQL_PORT", "30001"),
		MysqlDBName:      GetString("MYSQL_NAME", "balances"),
		AwsAddress:       GetString("AWS_ADDRESS", "http://192.168.49.2:30002"),
		AwsRegion:        GetString("AWS_REGION", "us-east-1"),
		AccessKeyID:      GetString("AWS_ACCESS_KEY_ID", "test"),
		SecretAccessKey:  GetString("AWS_SECRET_ACCESS_KEY", "test"),
		BalancesSNSTopic: GetString("BALANCES_SNS_TOPIC", "arn:aws:sns:us-east-1:000000000000:balances-sns-topic"),
		BalancesSQSQueue: GetString("BALANCES_SQS_QUEUE", "http://192.168.49.2:30002/000000000000/balances-sqs-queue"),
	}

	for _, optFunc := range options {
		optFunc(c)
	}

	return c
}

func WithAppPort(appPort int) Option {
	return func(c *Config) {
		c.AppPort = appPort
	}
}

func WithMySqlDBUser(mySqlDBUser string) Option {
	return func(c *Config) {
		c.MySqlDBUser = mySqlDBUser
	}
}

func WithMySqlDBPass(MySqlDBPass string) Option {
	return func(c *Config) {
		c.MySqlDBPass = MySqlDBPass
	}
}

func WithMySqlDBHost(MySqlDBHost string) Option {
	return func(c *Config) {
		c.MySqlDBHost = MySqlDBHost
	}
}
func WithMySqlDBPort(MySqlDBPort string) Option {
	return func(c *Config) {
		c.MySqlDBPort = MySqlDBPort
	}
}
func WithMysqlDBName(MysqlDBName string) Option {
	return func(c *Config) {
		c.MysqlDBName = MysqlDBName
	}
}
func WithAwsAddress(AwsAddress string) Option {
	return func(c *Config) {
		c.AwsAddress = AwsAddress
	}
}
func WithAwsRegion(AwsRegion string) Option {
	return func(c *Config) {
		c.AwsRegion = AwsRegion
	}
}

func WithSnsAccountTopic(SnsAccountTopic string) Option {
	return func(c *Config) {
		c.BalancesSNSTopic = SnsAccountTopic
	}
}

func (c Config) GetMySQLConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.MySqlDBUser,
		c.MySqlDBPass,
		c.MySqlDBHost,
		c.MySqlDBPort,
		c.MysqlDBName,
	)
}

func GetString(env string, def string) string {
	if e := os.Getenv(env); e != "" {
		return e
	}
	return def
}

func GetInt(env string, def int) int {
	i, err := strconv.Atoi(os.Getenv(env))
	if err != nil {
		return def
	}
	return i
}
