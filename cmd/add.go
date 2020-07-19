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

	"github.com/mrityunjaygr8/go-pass/stuff"
	"github.com/spf13/cobra"
)

var addURL string
var addUsername string
var addPassword string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the password for the given URL and username",
	Long:  `Add the password for the given URL and username`,
	Run: func(cmd *cobra.Command, args []string) {
		if addURL == "" || addUsername == "" || addPassword == "" {
			fmt.Println("Must include both URL, username and password")
			os.Exit(1)
		}

		s := initDB()
		i := stuff.Item{URL: addURL, Username: addUsername, Password: addPassword}
		err := s.AddItem(i)
		if err != nil {
			log.Fatal()
		}
		s.Close()
	},
}

func initDB() stuff.Store {
	s, err := stuff.Init()
	if err != nil {
		log.Fatal()
	}
	return s
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().StringVar(&addURL, "url", "", "the URL for which to add the password")
	addCmd.MarkFlagRequired("url")
	addCmd.PersistentFlags().StringVar(&addUsername, "username", "", "the username for which to add the password")
	addCmd.MarkFlagRequired("username")
	addCmd.PersistentFlags().StringVar(&addPassword, "password", "", "the password to add")
	addCmd.MarkFlagRequired("password")
}
