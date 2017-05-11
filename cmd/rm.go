package cmd

import (
	"github.com/sh3rp/s3/aws"
	"github.com/spf13/cobra"
)

var RmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove files in a bucket",
	Long:  "Remove files in a bucket",
	Run: func(cmd *cobra.Command, args []string) {
		svc := aws.GetService(region, bucket)
		var key string
		if len(args) > 2 {
			key = args[len(args)-1]
		} else {
			key = ""
		}
		files := svc.S3GetObjects(key)

		for _, f := range files {
			svc.S3RemoveObject(*f.Key)
		}
	},
}
