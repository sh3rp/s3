package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sh3rp/s3/aws"
	"github.com/spf13/cobra"
)

var PutCmd = &cobra.Command{
	Use:   "put",
	Short: "Put a file in a bucket",
	Long:  "Put a file in a bucket",
	Run: func(cmd *cobra.Command, args []string) {
		svc := aws.GetService(region, bucket)
		key := args[len(args)-1]
		files, err := get_all_files(key)

		if err != nil {
			fmt.Printf("Error uploading: %v\n", err)
		}

		for _, file := range files {
			svc.S3PutObject(file)
		}
	},
}

func get_all_files(filename string) ([]string, error) {
	var files []string

	if filename != "" && filename[len(filename)-1] == '*' {

	} else {
		info, err := os.Stat(filename)

		if err != nil {
			return nil, err
		}

		if info.IsDir() {
			files = get_all_files_in_directory(info.Name())
		} else {
			files = append(files, info.Name())
		}

	}

	return files, nil
}

func get_all_files_in_directory(dirname string) []string {
	var files []string

	filelist, err := ioutil.ReadDir(dirname)

	if err != nil {
		return nil
	}

	for _, file := range filelist {
		if file.IsDir() {
			files = append(files, get_all_files_in_directory(dirname+"/"+file.Name())...)
		} else {
			files = append(files, dirname+"/"+file.Name())
		}
	}

	return files
}
