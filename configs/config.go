package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// app port
	Port = 5000

	//mysql connection
	ConnectionString = ""

	// aws configurations
	AwsAddress      = ""
	AwsRegion       = ""
	AwsProfile      = ""
	AwsID           = ""
	AwsSecret       = ""
	SnsAccountTopic = ""
)

func Load() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ConnectionString = getMySQLConnectionString()

	AwsAddress = os.Getenv("AWS_ADDRESS")
	AwsRegion = os.Getenv("AWS_REGION")
	AwsProfile = os.Getenv("AWS_PROFILE")
	AwsID = os.Getenv("AWS_ID")
	AwsSecret = os.Getenv("AWS_SECRET")

	SnsAccountTopic = os.Getenv("ACCOUNT_SNS_TOPIC")
}

func getMySQLConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_DB_USUARIO"),
		os.Getenv("MYSQL_DB_SENHA"),
		os.Getenv("MYSQL_DB_HOST"),
		os.Getenv("MYSQL_DB_PORT"),
		os.Getenv("MYSQL_DB_NOME"),
	)
}
