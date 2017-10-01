package main

import (
	"fmt"
	"os"

	"github.com/sh3rp/s3/aws"
	"github.com/sh3rp/s3/cmd"
)

func main() {
	access_key := os.Getenv(aws.ACCESS_KEY)
	secret_access_key := os.Getenv(aws.SECRET_ACCESS_KEY)

	if access_key == "" {
		fmt.Printf("Please set your %s environment variable.\n", aws.ACCESS_KEY)
		os.Exit(1)
	}

	if secret_access_key == "" {
		fmt.Printf("Please set your %s environment variable.\n", aws.SECRET_ACCESS_KEY)
		os.Exit(1)
	}

	cmd.Init()
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
