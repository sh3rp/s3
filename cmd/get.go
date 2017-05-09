package cmd

import (
	"strings"

	"github.com/sh3rp/s3/aws"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a file in a bucket",
	Long:  "Get a file in a bucket",
	Run: func(cmd *cobra.Command, args []string) {
		svc := aws.GetService(region, bucket)
		if tofile == "" {
			tofile = key[strings.LastIndex(key, "/")+1:]
		}
		svc.S3GetObject(key, tofile)
	},
}
