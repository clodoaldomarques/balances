package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort         int
	MySqlDBUser     string
	MySqlDBPass     string
	MySqlDBHost     string
	MySqlDBPort     string
	MysqlDBName     string
	AwsAddress      string
	AwsRegion       string
	AwsProfile      string
	AwsID           string
	AwsSecret       string
	SnsAccountTopic string
}

type Option func(*Config)

func New(options ...Option) *Config {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	c := &Config{
		AppPort:         getInt("APP_PORT", 5000),
		MySqlDBUser:     getString("MYSQL_DB_USER", "admin"),
		MySqlDBPass:     getString("MYSQL_DB_PASS", "b4l4nc3s"),
		MySqlDBHost:     getString("MYSQL_DB_HOST", "192.168.49.2"),
		MySqlDBPort:     getString("MYSQL_DB_PORT", "30306"),
		MysqlDBName:     getString("MYSQL_DB_NAME", "balances"),
		AwsAddress:      getString("AWS_ADDRESS", "192.168.49.2"),
		AwsRegion:       getString("AWS_REGION", "us-east-1"),
		AwsProfile:      getString("AWS_PROFILE", "localstack"),
		AwsID:           getString("AWS_ID", "1.0.0"),
		AwsSecret:       getString("AWS_SECRET", "test"),
		SnsAccountTopic: getString("ACCOUNT_SNS_TOPIC", "arn:aws:sns:us-east-1:000000000000:account-sns-topic"),
	}

	for _, opt := range options {
		opt(c)
	}

	return c
}

func (c Config) WithAppPort(appPort int) Option {
	return func(c *Config) {
		c.AppPort = appPort
	}
}

func (c Config) WithMySqlDBUser(mySqlDBUser string) Option {
	return func(c *Config) {
		c.MySqlDBUser = mySqlDBUser
	}
}

func (c Config) WithMySqlDBPass(MySqlDBPass string) Option {
	return func(c *Config) {
		c.MySqlDBPass = MySqlDBPass
	}
}

func (c Config) WithMySqlDBHost(MySqlDBHost string) Option {
	return func(c *Config) {
		c.MySqlDBHost = MySqlDBHost
	}
}
func (c Config) WithMySqlDBPort(MySqlDBPort string) Option {
	return func(c *Config) {
		c.MySqlDBPort = MySqlDBPort
	}
}
func (c Config) WithMysqlDBName(MysqlDBName string) Option {
	return func(c *Config) {
		c.MysqlDBName = MysqlDBName
	}
}
func (c Config) WithAwsAddress(AwsAddress string) Option {
	return func(c *Config) {
		c.AwsAddress = AwsAddress
	}
}
func (c Config) WithAwsRegion(AwsRegion string) Option {
	return func(c *Config) {
		c.AwsRegion = AwsRegion
	}
}
func (c Config) WithAwsProfile(AwsProfile string) Option {
	return func(c *Config) {
		c.AwsProfile = AwsProfile
	}
}

func (c Config) WithAwsID(AwsID string) Option {
	return func(c *Config) {
		c.AwsID = AwsID
	}
}

func (c Config) WithAwsSecret(AwsSecret string) Option {
	return func(c *Config) {
		c.AwsSecret = AwsSecret
	}
}

func (c Config) WithSnsAccountTopic(SnsAccountTopic string) Option {
	return func(c *Config) {
		c.SnsAccountTopic = SnsAccountTopic
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

func getString(env string, def string) string {
	if e := os.Getenv(env); e != "" {
		return e
	}
	return def
}

func getInt(env string, def int) int {
	i, err := strconv.Atoi(os.Getenv(env))
	if err != nil {
		return def
	}
	return i
}
