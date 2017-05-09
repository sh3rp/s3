package cmd

import (
	"github.com/sh3rp/s3/aws"
	"github.com/spf13/cobra"
)

var PutCmd = &cobra.Command{
	Use:   "put",
	Short: "Put a file in a bucket",
	Long:  "Put a file in a bucket",
	Run: func(cmd *cobra.Command, args []string) {
		svc := aws.GetService(region, bucket)
		svc.S3PutObject(fromfile)
	},
}
