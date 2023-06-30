package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func setScanPersistentFlags(flags *pflag.FlagSet) {
	cwd, _ := os.Getwd()

	flags.String(
		"directory",
		cwd,
		"IaC root directory",
	)
	flags.String(
		"id",
		"",
		"Application's Client ID",
	)
}
func NewScanCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Scan client IaC with checkov",
		Run:   scanRun,
	}
	setScanPersistentFlags(cmd.PersistentFlags())
	return cmd
}

func scanRun(cmd *cobra.Command, args []string) {
	directory, err := cmd.Flags().GetString("directory")
	if err != nil {
		log.Fatal(err)
	}

	policies, err := getPolicies()
	if err != nil {
		log.Fatal(err)
	}

	policyNames := make([]string, 0, len(policies))
	for _, p := range policies {
		policyNames = append(policyNames, p.Name)
	}

	checkov := exec.Command(
		"checkov",
		"--directory",
		directory,
		"--check",
		strings.Join(policyNames, ","),
	)

	stdout, err := checkov.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	defer stdout.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	err = checkov.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = checkov.Wait()
	if err != nil {
		log.Fatal("something is wrong with checkov: " + err.Error())
	}
	wg.Wait()
}

type PoliciesResponse struct {
	Policies []Policy `json:"policies"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	TransactionID string `json:"transaction_id"`
}

type Policy struct {
	Name     string `json:"name,omitempty"`
	Version  string `json:"version,omitempty"`
	Location string `json:"location,omitempty"`
}

func getPolicies() ([]Policy, error) {
	var protocol = "http"
	var host = "localhost"
	var port = "8080"
	var slug = "policies"
	var version = "v1"

	client := http.Client{
		Timeout: time.Second * 3,
	}

	res, err := client.Get(
		fmt.Sprintf("%s://%s:%s/%s/%s", protocol, host, port, version, slug))
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	policiesResponse := PoliciesResponse{}
	err = json.Unmarshal(resBody, &policiesResponse)
	if err != nil {
		return nil, err
	}

	return policiesResponse.Policies, nil
}
