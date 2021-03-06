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

	"cer/config"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var vCmd = &cobra.Command{
	Use:   "v",
	Short: "show version info",
	Long: `BuildTime
	 BuildVersion
	 Commit hash`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version called")
		fmt.Println("Build Version:\t", config.BuildVersion)
		fmt.Println("Build Time:\t", config.BuildTime)
		fmt.Println("Commit hash:\t", config.CommitHash)

	},
}

func init() {
	rootCmd.AddCommand(vCmd)
}
