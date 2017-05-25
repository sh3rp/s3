package cmd

import (
	"fmt"

	"github.com/sh3rp/s3/aws"
	"github.com/spf13/cobra"
)

var CatCmd = &cobra.Command{
	Use:   "cat",
	Short: "Cat a file in a bucket",
	Long:  "Cat a file in a bucket",
	Run: func(cmd *cobra.Command, args []string) {
		svc := aws.GetService(region, bucket)
		if len(args) > 0 {
			key := args[len(args)-1]
			svc.S3CatObject(key)
		} else {
			fmt.Println("You must supply a file.")
		}
	},
}
