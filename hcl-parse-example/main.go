package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

type Provider struct {
	AWS AWSProvider `hcl:"provider,block"`
}

type AWSProvider struct {
	AWS        string          `hcl:"aws,label"`
	AccessKey  string          `hcl:"access_key"`
	SecretKey  string          `hcl:"secret_key"`
	Region     string          `hcl:"region"`
	Token      string          `hcl:"token"`
	AssumeRole AssumeRoleBlock `hcl:"assume_role,block"`
}

type AssumeRoleBlock struct {
	RoleARN    string `hcl:"role_arn"`
	ExternalID string `hcl:"external_id"`
}

func main() {
	pvdr := Provider{}
	parser := hclparse.NewParser()
	file, diag := parser.ParseHCLFile("main.tf")
	if diag.HasErrors() {
		log.Fatal(diag.Error())
	}
	diag = gohcl.DecodeBody(file.Body, nil, &pvdr)
	if diag.HasErrors() {
		log.Fatal(diag.Error())
	}
	fmt.Printf("%#v\n", pvdr)
}
