provider "aws" {
  access_key = "anyvalue"
  secret_key = "anyvalue"
  region = "us-east-1"
  s3_use_path_style           = true
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
  
  endpoints {
    apigateway     = "http://192.168.49.2:30566"
    apigatewayv2   = "http://192.168.49.2:30566"
    cloudformation = "http://192.168.49.2:30566"
    cloudwatch     = "http://192.168.49.2:30566"
    dynamodb       = "http://192.168.49.2:30566"
    ec2            = "http://192.168.49.2:30566"
    es             = "http://192.168.49.2:30566"
    elasticache    = "http://192.168.49.2:30566"
    firehose       = "http://192.168.49.2:30566"
    iam            = "http://192.168.49.2:30566"
    kinesis        = "http://192.168.49.2:30566"
    lambda         = "http://192.168.49.2:30566"
    rds            = "http://192.168.49.2:30566"
    redshift       = "http://192.168.49.2:30566"
    route53        = "http://192.168.49.2:30566"
    s3             = "http://s3.192.168.49.2.localstack.cloud:30566"
    secretsmanager = "http://192.168.49.2:30566"
    ses            = "http://192.168.49.2:30566"
    sns            = "http://192.168.49.2:30566"
    sqs            = "http://192.168.49.2:30566"
    ssm            = "http://192.168.49.2:30566"
    stepfunctions  = "http://192.168.49.2:30566"
    sts            = "http://192.168.49.2:30566"
  }

}