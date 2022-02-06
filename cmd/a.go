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
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// mCmd represents the m command
var aCmd = &cobra.Command{
	Use:   "a",
	Short: "CCOMs issue list",
	Long:  ` All CCOMs issue list in gitlab, according to related issue label`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("m called")
		repo_url := viper.GetString("project.repo_url")
		endp := "api/v4/projects/5674/issues?"
		repo := repo_url + endp
		c := internal.HttpClient()

		type User struct {
			Id string 
			Name string
		}
		n := []User{
			{ Id: "a00264776", Name: "Ahmet Duman", },
			{ Id: "a00537946", Name: "Ahmet Erol", },
			{ Id: "b00241020", Name: "Baris Can Menekse", },
			{ Id: "b00471095", Name: "Bora Gunyel", },
			{ Id: "b00543550", Name: "Burcu Bag", },
			{ Id: "b84199544", Name: "Burak Topcu", },
			{ Id: "c00600060", Name: "Cihan Biber", },
			{ Id: "c84191964", Name: "Can Uzunay", },
			{ Id: "f00598599", Name: "Fatih Dagli", },
			{ Id: "g84234118", Name: "Gamze Acil", },
			{ Id: "h84234119", Name: "Hikmet Cakir", },
			{ Id: "i00639667", Name: "Ibrahim Sahin", },
			{ Id: "i00747267", Name: "Ismail Oguz Saylan", },
			{ Id: "k00540961", Name: "Kubilay Ozata", },
			{ Id: "m00483517", Name: "Mustafa Erbay", },
			{ Id: "m84199810", Name: "Murat Yuksel", },
			{ Id: "m84203416", Name: "Murat Dogan", },
			{ Id: "n84199803", Name: "Neslihan Keser", },
			{ Id: "r84234114", Name: "Ramazan Biyik", },
			{ Id: "s00548611", Name: "Selim Sahin", },
			{ Id: "s00601060", Name: "Suha Kopan", },
			{ Id: "s00618143", Name: "Sevcan Erdogan", },
			{ Id: "u84184055", Name: "Umut Akdenizli", },
			{ Id: "wwx596396", Name: "wangpeixi", },
		}
		for _, v := range n{
			internal.GetIssuesByName(c, http.MethodGet, repo, "100", v.Id, v.Name)
		}
			
		// ch1 := make(chan bool)
		// go internal.GetIssuesByName(c, http.MethodGet, repo, "100", "m00483517", "Mustafa Erbay")
		// go internal.GetIssuesByName(c, http.MethodGet, repo, "100", "m00483517", "Mustafa Erbay" ,ch1)



	},
}

func init() {
	rootCmd.AddCommand(aCmd)

}
