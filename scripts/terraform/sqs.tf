resource "aws_sqs_queue" "balances-sqs-queue" {
  name = "balances-sqs-queue"
}

resource "aws_sns_topic_subscription" "balances-sns-sqs-subscription" {
  topic_arn = aws_sns_topic.balances-sns-topic.arn
  protocol  = "sqs"
  endpoint  = aws_sqs_queue.balances-sqs-queue.arn 
}