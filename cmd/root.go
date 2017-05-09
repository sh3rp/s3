package cmd

import "github.com/spf13/cobra"

var bucket string
var region string
var key string

var RootCmd = &cobra.Command{
	Use:   "s3",
	Short: "s3 utility",
	Long:  "Utility to work with Amazon AWS S3",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Init() {
	RootCmd.PersistentFlags().StringVarP(&region, "region", "r", "us-east-1", "AWS region to use")
	RootCmd.PersistentFlags().StringVarP(&bucket, "bucket", "b", "sh3rp", "AWS S3 bucket to use")
	RootCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "AWS S3 object key to use")
	RootCmd.AddCommand(LsCmd)
	RootCmd.AddCommand(GetCmd)
}
