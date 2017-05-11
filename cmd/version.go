package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version info",
	Long:  "Version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("s3 - v%d.%d\n", VERSION_MAJOR, VERSION_MINOR)
	},
}
