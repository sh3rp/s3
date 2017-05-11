package cmd

import (
	"fmt"

	"github.com/sh3rp/s3/aws"
	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Configuration info being used",
	Long:  "Configuration info being used",
	Run: func(cmd *cobra.Command, args []string) {
		svc := aws.GetService(region, bucket)
		fmt.Printf("\ns3 util version %d.%d\n\n", VERSION_MAJOR, VERSION_MINOR)
		fmt.Printf("Access key : %s\n", svc.Access_key)
		fmt.Printf("Secret key : %s\n", svc.Secret_access_key)
		fmt.Printf("Bucket     : %s\n", svc.Bucket)
		fmt.Printf("Region     : %s\n\n", svc.Region)
	},
}
