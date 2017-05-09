package main

import (
	"fmt"
	"os"

	"github.com/sh3rp/s3/cmd"
)

func main() {
	cmd.Init()
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
