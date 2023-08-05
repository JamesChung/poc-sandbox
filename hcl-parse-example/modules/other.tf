resource "aws_apigatewayv2_api" "example" {
  name                       = "example-websocket-api"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}

module "terraform_test_module" {
  source  = "spacelift.io/your-org/terraform_test_module"
  version = "1.0.0"
}

module "terraform_test_module2" {
  source  = "spacelift.io/your-org/terraform_test_module"
  version = "1.0.0"
}

provider "aws" {
  alias  = "west"
  region = "us-west-2"
  assume_role {
    # The role ARN within Account B to AssumeRole into. Created in step 1.
    role_arn = "arn:aws:iam::01234567890:role/role_in_account_a"
    # (Optional) The external ID created in step 1c.
    external_id = "my_external_id"
  }
}

resource "aws_apigatewayv2_api" "example2" {
  provider                   = aws.west
  name                       = "example-websocket-api"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}
