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
	"fmt"
	"log"
	"net/http"
	_ "strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xuri/excelize/v2"
)
var (
	listOfUsers string
)
type User struct {
	Id string 
	Name string
}
// dCmd represents the d command
var dCmd = &cobra.Command{
	Use:   "d",
	Short: "A brief description of your command",
	Long: `User listesini excelden okuyarak yapilan bir ornek bu`,
	// Args : cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("d called")
		repo_url := viper.GetString("project.repo_url")
		endp := "api/v4/projects/5674/issues?"
		repo := repo_url + endp
		c := internal.HttpClient()
		f, err := excelize.OpenFile(listOfUsers)
		if err != nil {
			log.Fatal(err)
		}
		
		ids, _ := internal.GetCol(f,"Sheet1",'A')
		names, _ := internal.GetCol(f,"Sheet1",'B')
		if len(ids) != len(names) {
			log.Fatal("user id and username does not match")
		}
		idList := make([]string, 0, len(ids))
		nameList := make([]string, 0, len(names))
		idList = append(idList, ids...)
		nameList = append(nameList, names...)

		fmt.Println(len(nameList))
		fmt.Println(idList)
		var userList []User
		for i := 0; i < len(idList); i++ {
			userList = append(userList,User{
				Id: idList[i],
				Name: nameList[i],
			})
		}
		fmt.Println(userList[1].Id)
		for i := range userList{
			internal.GetIssuesByName(c, http.MethodGet, repo, "100", userList[i].Id, userList[i].Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(dCmd)
	flags := dCmd.Flags()
	flags.StringVarP(&listOfUsers, "listOfUsers", "s", "", "gitlab user list required as sample.xlsx format")
	dCmd.MarkFlagRequired("listOfUsers")
	// dCmd.Flags()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flags.StringVarP(&listOfUsers,"listOfUsers", "s","", "A help for foo")
	// flags.StringVar(&listOfUsers, "listOfUsers", "", "listOfUsers")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	
}
