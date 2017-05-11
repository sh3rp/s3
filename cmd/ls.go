package cmd

import (
	"github.com/sh3rp/s3/aws"
	"github.com/spf13/cobra"
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all files in a bucket",
	Long:  "List all files in a bucket",
	Run: func(cmd *cobra.Command, args []string) {
		svc := aws.GetService(region, bucket)
		var key string
		if len(args) > 2 {
			key = args[len(args)-1]
		} else {
			key = ""
		}
		svc.S3ListBucket(key)
	},
}
