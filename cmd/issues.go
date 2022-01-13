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
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

// type User struct {
// 	Username string `json:"username"`
// 	Active bool `json:"active"`
// 	Id string `json:"id"`
// }
// issuesCmd represents the issues command
var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "issue list",
	Long: `Get issues with given filters`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("issues called")
		yourtokengoeshere := viper.GetString("personal_access_token")
		url := viper.GetString("repo_url")
		git, err := gitlab.NewClient(
			yourtokengoeshere, 
			gitlab.WithBaseURL(url),
		)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
		// newUser := User{
		// 	Active: true,
		// 	Username: "mustafa erbay",
		// }
		// users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{
		// 	Active: &newUser.Active,
		// 	Username: &newUser.Username,
			
		// })
		// issues, _, err := git.Issues.ListIssues(&gitlab.ListIssuesOptions{
		// 	AssigneeUsername: utils.StringPtr("m00483517"),
		// })

		projects, _, err := git.Projects.ListProjects(nil)
		if err != nil {
			log.Fatal(err)
		}
	
		log.Printf("Found %d projects", len(projects))

		// fmt.Println("user Count:", len(users))
		// if err != nil {
		// 	fmt.Errorf("Failed to list users: %v", err)
		// }
		for _, v := range projects {
		 	fmt.Println(v)
		}
	},
}

func init() {
	getCmd.AddCommand(issuesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// issuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// issuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
