package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Some scanner",
		Run:   runScan,
	}
	return cmd
}

func runScan(cmd *cobra.Command, args []string) {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(""))
	if err != nil {
		log.Fatal(err)
	}

	client := sts.NewFromConfig(cfg)

	identity, err := client.GetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ARN:", aws.ToString(identity.Arn))
	fmt.Println("Account:", aws.ToString(identity.Account))
	fmt.Println("UserID:", aws.ToString(identity.UserId))
	v, err := arn.Parse(aws.ToString(identity.Arn))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", v)
	// We probably want to print to the user what their scan ARN is

	// Partition Key:
	// USERID#<Some ARN>
	// Sort Key:
	// SCANID#<Some ULID>
}
