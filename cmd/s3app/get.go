package s3app

import (
	"fmt"

	"github.com/benmcrae/cmdapp/cmd/pkg/cloudprovider"
	"github.com/benmcrae/cmdapp/cmd/pkg/prompts"
	"github.com/manifoldco/promptui"
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

	s3Objects := awsClient.GetS3Objects(selBucket)

	selected := selectObjects(s3Objects)

	fmt.Println(selected)
}

func selectObjects(objects []cloudprovider.S3Object) string {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0000276F {{ .ObjectName | cyan }} ({{ .SizeMB | red }})",
		Inactive: "  {{ .ObjectName | cyan }} ({{ .SizeMB | red }})",
		Selected: "\U00002713 {{ .ObjectName | red | cyan }}",
		Details: `
--------- Accounts Details ----------
{{ "Name:" | faint }}	{{ .ObjectName }}
{{ "Size MB:" | faint }}	{{ .SizeMB }}
{{ "Last Modified:" | faint }}	{{ .LastModified }}
{{ "Storage Class:" | faint }}	{{ .StorageClass }}`,
	}

	prompt := promptui.Select{
		Label:     "S3 Objects",
		Items:     objects,
		Templates: templates,
		Size:      10,
	}

	_, selObject, _ := prompt.Run()

	return selObject

}
