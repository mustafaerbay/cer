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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// "github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

// tCmd represents the t command
var tCmd = &cobra.Command{
	Use:   "t",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("t called")
		_ , b := os.LookupEnv("GITLAB_USER_PASSWORD")
		if  !b {
			log.Fatal("GITLAB_USER_PASSWORD not exist in env variables")
		}
		client, err := gitlab.NewBasicAuthClient(
			viper.GetString("project.username"),
			os.Getenv("GITLAB_USER_PASSWORD"),
			gitlab.WithBaseURL("https://rnd-gitlab-eu.huawei.com/"),
		)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}

		fmt.Println(client.Issues.ListProjectIssues(5674, &gitlab.ListProjectIssuesOptions{
			ListOptions: gitlab.ListOptions{
				PerPage: 2,
			},
			OrderBy: gitlab.String("due_date"),
		}))
		// fmt.Println(git.Issues.ListIssues(&gitlab.ListIssuesOptions{
		// 	ListOptions: gitlab.ListOptions{
		// 		PerPage: 1,
		// 	},
		// }))

	},
}

func init() {
	rootCmd.AddCommand(tCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	tCmd.PersistentFlags().String("repo_url", "-r", "A help for foo")
	tCmd.Flags()
}

// GitlabSCM implements the SCM interface.
type GitlabSCM struct {
	client *gitlab.Client
}

// NewGitlabSCMClient returns a new GitLab client implementing the SCM interface.
func NewGitlabSCMClient(token string) *GitlabSCM {
	cli, _ := gitlab.NewOAuthClient(token, gitlab.WithoutRetries())
	return &GitlabSCM{
		client: cli,
	}
}
