resource "aws_sns_topic" "balances-sns-topic" {
  name = "balances-sns-topic"
}

resource "aws_sqs_queue" "balances-sqs" {
  name = "balances-sqs"
}

resource "aws_sns_topic_subscription" "balances-sns-sqs-subscription" {
  topic_arn = aws_sns_topic.balances-sns-topic.arn
  protocol  = "sqs"
  endpoint  = aws_sqs_queue.balances-sqs.arn 
}


resource "aws_dynamodb_table" "accounts" {
  name           = "accounts"
  billing_mode   = "PROVISIONED"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "account_id"
  range_key      = "org_id"

  attribute {
    name = "account_id"
    type = "N"
  }

  attribute {
    name = "org_id"
    type = "S"
  }

  tags = {
    Name        = "accounts"
    Environment = "production"
  }
}

resource "aws_dynamodb_table" "entries" {
  name           = "entries"
  billing_mode   = "PROVISIONED"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "tracking_id"
  range_key      = "account_id"

  attribute {
    name = "tracking_id"
    type = "S"
  }

  attribute {
    name = "account_id"
    type = "N"
  }

  tags = {
    Name        = "entries"
    Environment = "production"
  }
}