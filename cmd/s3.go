package cmd

import (
	"github.com/benmcrae/cmdapp/cmd/s3app"
	"github.com/spf13/cobra"
)

// s3Cmd represents the s3 command
var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "A brief description of your command",
}

func init() {
	s3Cmd.AddCommand(
		s3app.S3List(),
	)
}
