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
	AwsAddress              = ""
	AwsRegion               = ""
	AwsProfile              = ""
	AwsID                   = ""
	AwsSecret               = ""
	SnsCreateAccountTopic   = ""
	SnsChangeAvailableTopic = ""
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

	SnsCreateAccountTopic = os.Getenv("CREATE_ACCOUNT_SNS_TOPIC")
	SnsChangeAvailableTopic = os.Getenv("CHANGE_AVAILABLE_SNS_TOPIC")
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

func getMongoDBConnectionURL() string {
	database := os.Getenv("MONGO_DB_NOME")
	user := os.Getenv("MONGO_DB_USUARIO")
	pass := os.Getenv("MONGO_DB_SENHA")
	host := os.Getenv("MONGO_DB_HOST")
	port := os.Getenv("MONGO_DB_PORT")

	if user == "" || pass == "" {
		return fmt.Sprintf("mongodb://%s:%s/%s?replicaSet=rs1", host, port, database)
		//	return fmt.Sprintf("mongodb://%s:%s/%s?directConnection=true", host, port, database)
	}

	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?directConnection=true&replicaSet=rs1", user, pass, host, port, database)
	//return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?directConnection=true", user, pass, host, port, database)
}
