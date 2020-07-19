/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getURL string
var getUsername string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the password for the given URL and username",
	Long:  `Get the password for the given URL and username`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("URL: %s\nUsername: %s\n", getURL, getUsername)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().StringVar(&getURL, "url", "", "the URL for which to retrieve the password")
	getCmd.MarkFlagRequired("url")
	getCmd.PersistentFlags().StringVar(&getUsername, "username", "", "the username for which to retrieve the password")
	getCmd.MarkFlagRequired("username")
}
