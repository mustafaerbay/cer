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
var mCmd = &cobra.Command{
	Use:   "m",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("m called")
		repo_url := viper.GetString("repo_url")
		endp := "api/v4/projects/5674/issues?"
		repo := repo_url + endp
		c := internal.HttpClient()

		internal.GetIssuesByName(c, http.MethodGet, repo, "100", "a00264776", "Ahmet Duman")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","a00537946","Ahmet Erol")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","a00608324","ATACAN KULLABCI")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","b00241020","Baris Can Menekse")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","b00471095","Bora Gunyel")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","b00543550","Burcu Bag")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","b84199544","Burak Topcu")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","c00600060","Cihan Biber")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","c84191964","Can Uzunay")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","f00598599","Fatih Dagli")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","g84234118","Gamze Acil")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","h84234119","Hikmet Cakir")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","i00639667","Ibrahim Sahin")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","i00747267","Ismail Oguz Saylan")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","k00540961","Kubilay Ozata")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","m00483517","Mustafa Erbay")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","m84199810","Murat Yuksel")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","m84203416","Murat Dogan")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","n84199803","Neslihan Keser")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","r84234114","Ramazan Biyik")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","s00324983","Sevket Yurdacan")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","s00548611","Selim Sahin")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","s00601060","Suha Kopan")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","s00618143","Sevcan Erdogan")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","u84184055","Umut Akdenizli")
		// internal.GetIssuesByName(c,http.MethodGet,repo,"100","wwx596396","wangpeixi")

	},
}

func init() {
	rootCmd.AddCommand(mCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
