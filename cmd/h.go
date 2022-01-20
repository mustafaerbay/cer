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
	"cer/internal"
	"cer/modals"
	"encoding/json"
	"net/http"

	_ "encoding/json"
	_ "log"

	"fmt"

	_ "github.com/Jeffail/gabs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// hCmd represents the h command

var hCmd = &cobra.Command{
	Use:   "h",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("h called")
		// https://gitlab.com/api/v4/issues?scope=all&state=opened&assignee_username=derekferguson&not[labels]=Category:DAST,devops::secure&not[milestone]=13.11
		repo_url := viper.GetString("repo_url")
		endp := "api/v4/projects/5674/issues?"
		repo := repo_url + endp
		as := "assignee_username=c00600060"
		// lbl := "labels=Type Bug&"
		pp := "per_page=100&"
		// defer profile.Start(profile.CPUProfile, profile.ProfilePath("./cpu")).Stop()
		// url := viper.GetString("repo_url") + "api/v4/projects/5674/issues?per_page=100"
		// url := internal.Joinstr(repo_url,endp,pp)
		url := repo + pp + as
		// url := viper.GetString("repo_url") + "api/v4/projects?per_page=25"
		fmt.Println(url)
		var issueBody []modals.Issue
		c := internal.HttpClient()
		responseBody := internal.SendRequest(c, http.MethodGet, url)
		json.Unmarshal(responseBody, &issueBody)
		fmt.Println("size:", len(issueBody))
		for _, v := range issueBody {
			// fmt.Println(v.IID, "|" ,v.Title)
			fmt.Println(v.IID, "|", v.Assignee.Name, "		|", v.DueDate)
		}

		internal.GetIssuesByName(c, http.MethodGet, repo, "100", "c00600060", "Cihan Biber")

	},
}

func init() {
	rootCmd.AddCommand(hCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
