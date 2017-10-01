package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VERSION_MAJOR = 1
var VERSION_MINOR = 0

var bucket string
var region string

var RootCmd = &cobra.Command{
	Use:   "s3",
	Short: "s3 utility",
	Long:  "Utility to work with Amazon AWS S3",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Type 's3 help' for a list of commands")
	},
}

func Init() {
	RootCmd.PersistentFlags().StringVarP(&region, "region", "r", "", "AWS region to use")
	RootCmd.PersistentFlags().StringVarP(&bucket, "bucket", "b", "", "AWS S3 bucket to use")
	RootCmd.AddCommand(LsCmd)
	RootCmd.AddCommand(GetCmd)
	RootCmd.AddCommand(PutCmd)
	RootCmd.AddCommand(VersionCmd)
	RootCmd.AddCommand(InfoCmd)
	RootCmd.AddCommand(RmCmd)
	RootCmd.AddCommand(TagCmd)
	RootCmd.AddCommand(CatCmd)
}
