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
}

resource "aws_apigatewayv2_api" "example2" {
  provider                   = aws.west
  name                       = "example-websocket-api"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}
