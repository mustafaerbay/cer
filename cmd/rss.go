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
	"github.com/mmcdole/gofeed"
	"github.com/spf13/cobra"
	"log"
)

type issue struct {
	title  string
	author string
}

// rssCmd represents the users command
var rssCmd = &cobra.Command{
	Use:   "rss",
	Short: "List gitlab users",
	Long:  "List gitlab users",
	Run: func(cmd *cobra.Command, args []string) {

		fp := gofeed.NewParser()
		gitlabURL := "https://rnd-gitlab-eu.huawei.com/htrdc-isd/ebg/CCOMS/-/issues.atom?feed_token=crmDdCRB8H_HwasY6WRS&state=opened&per_page=100"
		feed, err := fp.ParseURL(gitlabURL)
		checkError("gitlabURL not reachable", err)

		for i := 0; i < len(feed.Items); i++ {
			// fmt.Printf("%s|%s", feed.Items[i].Title, feed.Items[i].Author.Name)
			// 	fmt.Println("")
			newissue := issue{
				title:  feed.Items[i].Title,
				author: feed.Items[i].Author.Name,
			}
			// val1=s+strconv.Itoa(i+1)
			// val2string := strconv.Itoa(i+1) + "|" + newissue.title + "|" + newissue.author
			// fmt.Println(val2string)
			fmt.Println("-------------------------------------")
			fmt.Println(newissue.author)
			fmt.Println(newissue.title)

		}
	},
}

func init() {
	getCmd.AddCommand(rssCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rssCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rssCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func checkError(message string, err error) {
	if err != nil {
		log.Fatalln(message, err)
	}
}
