package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/perimeterx/marshmallow"
	"github.com/tmccombs/hcl2json/convert"
)

// GetUnifedJSONHCL returns a JSON representation of Terraform
// in a given repository
func GetUnifiedJSONTerraform(dir string) ([]byte, error) {
	if !path.IsAbs(dir) {
		return nil, fmt.Errorf("dir %s is not an absolute path", dir)
	}

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

func main() {
	fp, _ := filepath.Abs("./modules")
	b, err := GetUnifiedJSONTerraform(fp)
	if err != nil {
		log.Fatal(err)
	}

	result, err := marshmallow.Unmarshal(b, &struct{}{})
	if err != nil {
		log.Fatal(err)
	}

	val, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(val))

	// fmt.Printf("%#v\n", result)
	// fmt.Println("---")
	// fmt.Printf("%#v\n", result["module"])
	// m := result["module"].(map[string]any)
	// fmt.Println("------")
	// someVal, err := marshmallow.UnmarshalFromJSONMap(m, &struct{}{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(someVal)
	// fmt.Println("--")
	// for k, v := range someVal {
	// 	fmt.Println(k, v)
	// }
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
	if !path.IsAbs(dir) {
		return nil, fmt.Errorf("dir %s is not an absolute path", dir)
	}

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
