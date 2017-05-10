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
		key := args[len(args)-1]
		svc.S3ListBucket(key)
	},
}
