package s3app

import (
	"fmt"

	"github.com/benmcrae/cmdapp/cmd/pkg/cloudprovider"
	"github.com/spf13/cobra"
)

func S3List() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List S3 buckets",
		Run:   run,
	}

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	awsClient := cloudprovider.NewClients()

	resp := awsClient.ListS3Buckets()

	for _, b := range resp {
		fmt.Println(b.BucketName)
	}
}
