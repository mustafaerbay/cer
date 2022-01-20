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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type bpi struct {
	Code        string `json:"code"`
	Rate        string `json:"rate"`
	Description string `json:"description"`
}

// dolarCmd represents the dolar command
var dolarCmd = &cobra.Command{
	Use:   "dolar",
	Short: "1 dolar kac TL",
	Long: `1 dolarin TL karsiligini vermeye yarar 
	
	Isterseniz asagidaki sekilde carpa islemi de yaptirabilirsiniz
	
	cer.exe dolar -a 100 
	100 dolarin TL karsiligini vereecektir`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://api.coindesk.com/v1/bpi/currentprice.json"
		res, err := http.Get(url)

		if err != nil {
			panic(err.Error())
		}

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			panic(err.Error())
		}

		// var data bpi
		data2, _ := json.Marshal(body)

		fmt.Printf("Results: %v\n", data2)
		os.Exit(0)

	},
}

func init() {
	rootCmd.AddCommand(dolarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dolarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dolarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
