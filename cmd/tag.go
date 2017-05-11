package cmd

import (
	"fmt"
	"strings"

	"github.com/sh3rp/s3/aws"
	"github.com/spf13/cobra"
)

var TagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Tag an object in a bucket",
	Long:  "Tag an object in a bucket",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Printf("Must supply a key and tags\n")
			return
		}
		svc := aws.GetService(region, bucket)
		tags := make(map[string]string)
		key := args[0]
		tagsList := args[1]

		kvList := strings.Split(tagsList, ",")

		for _, kv := range kvList {
			keyValue := strings.Split(kv, "=")
			tags[keyValue[0]] = keyValue[1]
		}

		svc.S3TagObject(key, tags)
	},
}
