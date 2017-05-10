package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sh3rp/s3/aws"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a file in a bucket",
	Long:  "Get a file in a bucket",
	Run: func(cmd *cobra.Command, args []string) {
		svc := aws.GetService(region, bucket)
		key := args[len(args)-1]
		objects := svc.S3GetObjects(key)

		if len(objects) < 2 {
			fmt.Printf("Retrieving %s...\n", key)
			svc.S3GetObject(key)
		} else {
			get_all_objects(svc, objects)
		}
	},
}

func get_all_objects(svc *aws.S3Service, objects []*s3.Object) {
	for _, obj := range objects {
		fmt.Printf("Retrieving %s...\n", *obj.Key)
		svc.S3GetObject(*obj.Key)
	}
}
