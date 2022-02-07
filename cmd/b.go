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
var searchProject string
// bCmd represents the issues command
var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Project list by membership",
	Long:  `Get issues with given filters`,
	// Args : cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("b called")
		yourtokengoeshere := viper.GetString("project.personal_access_token")
		url := viper.GetString("project.repo_url")
		git, err := gitlab.NewClient(
			yourtokengoeshere,
			gitlab.WithBaseURL(url),
		)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}

		projects, _, err := git.Projects.ListProjects(
			&gitlab.ListProjectsOptions{
				Membership: gitlab.Bool(true),
				ListOptions: gitlab.ListOptions{
					Page: 1,
					PerPage: 100,
				},
				Search: gitlab.String(searchProject),
			},
			nil,)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Found %d projects", len(projects))

		for _, v := range projects {
			fmt.Println(v.Name,"|",v.CreatorID,"|",v.HTTPURLToRepo)
			fmt.Println("--------------------------------")
		}
		// fmt.Print(projects[0])
		fmt.Println(len(projects))
	},
}

func init() {
	rootCmd.AddCommand(bCmd)
	bCmd.Flags().StringVarP(&searchProject, "searchProject", "s", "", "project search")
	bCmd.MarkFlagRequired("searchProject")
	// bCmd.Flags()
}
