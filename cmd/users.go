/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	// "os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
	"log"
)

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "List gitlab users",
	Long:  "List gitlab users",
	Run: func(cmd *cobra.Command, args []string) {

		yourtokengoeshere := viper.GetString("personal_access_token")
		url := viper.GetString("repo_url")
		git, err := gitlab.NewClient(yourtokengoeshere, gitlab.WithBaseURL(url))
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
		users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{
			Active:   gitlab.Bool(true),
			Username: gitlab.String("m00483517"),
		})
		if err != nil {
			fmt.Errorf("Failed to list users: %v", err)
		}
		fmt.Println("user Count:", len(users))
		for _, v := range users {
			fmt.Println(v)
		}
		get_users, _, err := git.Users.GetUserMemberships(*gitlab.Int(483517), &gitlab.GetUserMembershipOptions{})
		if err != nil {
			fmt.Errorf("Failed to GetUserMemberships users: %v", err)
		}

		for _, v := range get_users {
			fmt.Println(v)
		}

	},
}

func init() {
	getCmd.AddCommand(usersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
