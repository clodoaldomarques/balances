variable "account-id" {
  type = string
  default = "000000000000"
}

resource "aws_sns_topic" "balances-sns-topic" {
  name = "balances-sns-topic"
}

resource "aws_sns_topic_policy" "default" {
  arn = aws_sns_topic.balances-sns-topic.arn

  policy = data.aws_iam_policy_document.balances_sns_topic_policy.json
}

data "aws_iam_policy_document" "balances_sns_topic_policy" {
  policy_id = "__default_policy_ID"

  statement {
    actions = [
      "SNS:Subscribe",
      "SNS:SetTopicAttributes",
      "SNS:RemovePermission",
      "SNS:Receive",
      "SNS:Publish",
      "SNS:ListSubscriptionsByTopic",
      "SNS:GetTopicAttributes",
      "SNS:DeleteTopic",
      "SNS:AddPermission",
    ]

    condition {
      test     = "StringEquals"
      variable = "AWS:SourceOwner"

      values = [
        var.account-id,
      ]
    }

    effect = "Allow"

    principals {
      type        = "AWS"
      identifiers = ["*"]
    }

    resources = [
      aws_sns_topic.balances-sns-topic.arn,
    ]

    sid = "__default_statement_ID"
  }
}
