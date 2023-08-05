provider "aws" {
  # The security credentials for AWS Account A.
  access_key = "AKIAXXXXXXXXXXXXXXXX"
  secret_key = "123XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  region     = "us-east-1"
  where      = "modules"

  # (Optional) the MFA token for Account A.
  token = "123456"

  assume_role {
    # The role ARN within Account B to AssumeRole into. Created in step 1.
    role_arn = "arn:aws:iam::01234567890:role/role_in_account_b"
    # (Optional) The external ID created in step 1c.
    external_id = "my_external_id"
  }
}

provider "azure" {
  # The security credentials for AWS Account A.
  access_key = "AKIAXXXXXXXXXXXXXXXX"
  secret_key = "123XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  region     = "us-east-1"
  where      = "modules"

  # (Optional) the MFA token for Account A.
  token = "123456"

  assume_role {
    # The role ARN within Account B to AssumeRole into. Created in step 1.
    role_arn = "arn:aws:iam::01234567890:role/role_in_account_c"
    # (Optional) The external ID created in step 1c.
    external_id = "my_external_id"
  }
}
