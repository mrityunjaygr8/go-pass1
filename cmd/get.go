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
	"log"
	"os"

	"github.com/spf13/cobra"
)

var getURL string
var getUsername string
var getAllUsers bool
var getAllURLs bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the password for the given URL and username",
	Long:  `Get the password for the given URL and username`,
	Run: func(cmd *cobra.Command, args []string) {
		if getAllUsers {
			if getAllURLs {
				fmt.Println("You cannot use all-urls flag with the all-users flag")
				os.Exit(1)
			}
			if getURL == "" {
				fmt.Println("You must provide a URL with the all-users flag")
				os.Exit(1)
			}
			s := initDB()
			err := s.GetAllURLUsers(getURL)
			if err != nil {
				log.Fatal(err)
			}
			s.Close()
			os.Exit(0)
		}

		if getAllURLs {
			s := initDB()
			err := s.GetAllURLs()
			if err != nil {
				log.Fatal(err)
			}
			s.Close()
			os.Exit(0)
		}

		if getURL == "" || getUsername == "" {
			fmt.Println("Must include both URL and username")
			os.Exit(1)
		} else {
			s := initDB()
			err := s.GetItem(getURL, getUsername)
			if err != nil {
				log.Fatal(err)
			}
			s.Close()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().StringVar(&getURL, "url", "", "the URL for which to retrieve the password")
	getCmd.MarkFlagRequired("url")
	getCmd.PersistentFlags().StringVar(&getUsername, "username", "", "the username for which to retrieve the password")
	getCmd.MarkFlagRequired("username")
	getCmd.PersistentFlags().BoolVar(&getAllUsers, "all-users", false, "get all users for a given URL")
	getCmd.MarkFlagRequired("all-users")
	getCmd.PersistentFlags().BoolVar(&getAllURLs, "all-urls", false, "get all URLs for which credentials are stored")
	getCmd.MarkFlagRequired("all-urls")
}
