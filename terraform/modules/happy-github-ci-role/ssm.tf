data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "ssm_reader_writer" {
  statement {
    sid = "GhActionsSSMReaderWriter"
    actions = [
      "ssm:Get*",
      "ssm:Put*"
    ]
    resources = [
      # this is the legacy location of SSM parameters
      "arn:aws:ssm:us-west-2:${data.aws_caller_identity.current.account_id}:parameter/happy/${var.tags.env}/*",
      # this is the new location of SSM parameters (namespaced on the happy app)
      "arn:aws:ssm:us-west-2:${data.aws_caller_identity.current.account_id}:parameter/happy/${var.happy_app_name}/${var.tags.env}/*"
    ]
  }
}
resource "aws_iam_role_policy" "ssm_reader_writer" {
  name   = "gh_actions_ssm_reader_writer_${local.namespace}"
  policy = data.aws_iam_policy_document.ssm_reader_writer.json
  role   = local.role_name

  depends_on = [module.gh_actions_role]
}
