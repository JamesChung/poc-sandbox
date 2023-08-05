package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/tmccombs/hcl2json/convert"
)

// GetUnifedJSONHCL returns a JSON representation of Terraform
// in a given repository
func GetUnifiedJSONTerraform(dir string) ([]byte, error) {
	v, err := GetTerraformFiles(dir)
	if err != nil {
		return nil, err
	}

	fbs, err := ReadFiles(v)
	if err != nil {
		return nil, err
	}

	unifiedTerraform := FileBytesSlice(fbs).Concatenate()

	parser := hclparse.NewParser()

	f, d := parser.ParseHCL(unifiedTerraform, "unifiedTerraform.tf")
	if d.HasErrors() {
		log.Fatal(d.Error())
	}

	b, err := convert.File(f, convert.Options{})
	if err != nil {
		log.Fatal(err)
	}

	return b, nil
}

func GetUnifiedTerraform(dir string) (*hcl.File, error) {
	v, err := GetTerraformFiles(dir)
	if err != nil {
		return nil, err
	}

	fbs, err := ReadFiles(v)
	if err != nil {
		return nil, err
	}

	unifiedTerraform := FileBytesSlice(fbs).Concatenate()

	parser := hclparse.NewParser()

	f, d := parser.ParseHCL(unifiedTerraform, "unifiedTerraform.hcl")
	if d.HasErrors() {
		log.Fatal(d.Error())
	}

	return f, nil
}

func GetRoleARNsFromAWSProviders(f *hcl.File) ([]string, error) {
	var t Terraform

	diag := gohcl.DecodeBody(f.Body, nil, &t)
	if diag.HasErrors() {
		for _, e := range diag.Errs() {
			return nil, errors.Join(e)
		}
	}

	roleARNs := make([]string, 0, len(t.Providers))

	for _, v := range t.Providers {
		if v.Name != "aws" {
			continue
		}
		var as AssumeRole
		gohcl.DecodeBody(v.Remain, nil, &as)
		if len(as.Block.RoleARN) > 0 {
			roleARNs = append(roleARNs, as.Block.RoleARN)
		}
	}

	return roleARNs, nil
}

type Terraform struct {
	Providers []Provider `hcl:"provider,block"`
	Remain    hcl.Body   `hcl:",remain"`
}

type Provider struct {
	Name   string   `hcl:"name,label"`
	Alias  string   `hcl:"alias,optional"`
	Region string   `hcl:"region,optional"`
	Remain hcl.Body `hcl:",remain"`
}

type AssumeRole struct {
	Block  AssumeRoleBlock `hcl:"assume_role,block"`
	Remain hcl.Body        `hcl:",remain"`
}

type AssumeRoleBlock struct {
	RoleARN string   `hcl:"role_arn,optional"`
	Remain  hcl.Body `hcl:",remain"`
}

func main() {
	fp, err := filepath.Abs("./modules")
	if err != nil {
		log.Fatal(err)
	}

	h, err := GetUnifiedTerraform(fp)
	if err != nil {
		log.Fatal(err)
	}

	arns, err := GetRoleARNsFromAWSProviders(h)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(arns)
}

type FileBytesSlice [][]byte

func (fbs FileBytesSlice) Concatenate() []byte {
	catSlice := []byte{}

	for _, b := range fbs {
		catSlice = append(catSlice, b...)
	}

	return catSlice
}

func GetTerraformFiles(dir string) ([]string, error) {
	dirEntry, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Create regex matcher for .tf files
	re, err := regexp.Compile(`^.*\.(tf)$`)
	if err != nil {
		return nil, err
	}

	ss := []string{}

	// Find all .tf matches in directory
	for _, v := range dirEntry {
		if v.Type().IsRegular() {
			if re.MatchString(v.Name()) {
				ss = append(ss, filepath.Join(dir, v.Name()))
			}
		}
	}

	return ss, nil
}

func ReadFiles(files []string) ([][]byte, error) {
	fileBytesSlice := [][]byte{}
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		fileBytesSlice = append(fileBytesSlice, b)
	}
	return fileBytesSlice, nil
}
