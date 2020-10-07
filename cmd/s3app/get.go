package s3app

import (
	"fmt"

	"github.com/benmcrae/cmdapp/cmd/pkg/cloudprovider"
	"github.com/benmcrae/cmdapp/cmd/pkg/prompts"
	"github.com/spf13/cobra"
)

func S3Get() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get S3 object",
		Run:   runS3Get,
	}

	return cmd
}

func runS3Get(cmd *cobra.Command, args []string) {
	BucketNames := []string{}

	awsClient := cloudprovider.NewClients()
	resp := awsClient.ListS3Buckets()

	for _, i := range resp {
		BucketNames = append(BucketNames, i.BucketName)
	}

	selBucket, err := prompts.GetListValue(BucketNames, "Please select bucket")
	if err != nil {
		fmt.Errorf("Error getting selected list value %v", err.Error())
		return
	}

	fmt.Println(selBucket)
}
