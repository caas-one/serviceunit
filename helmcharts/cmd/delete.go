// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete命令调用helm删除所有的应用",
	Long:  `delete命令调用helm删除所有的应用`,
	Run: func(cmd *cobra.Command, args []string) {
		if deleteAll {
			deleteAllHelm(root, profile)
		} else if !strings.EqualFold(deleteSu, "") && strings.EqualFold(deleteApp, "") {
			deleteSuHelm(root, profile, deleteSu)
		} else if !strings.EqualFold(deleteSu, "") && !strings.EqualFold(deleteApp, "") {
			deleteAppHelm(root, profile, deleteSu, deleteApp)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	deleteCmd.PersistentFlags().StringVarP(&profile, "profile", "p", "profile.yaml", "profile文件的位置")
	deleteCmd.PersistentFlags().BoolVarP(&deleteAll, "all", "a", false, "安装所有的charts文件")
	deleteCmd.PersistentFlags().StringVarP(&deleteSu, "su", "s", "", "安装指定的服务单元")
	deleteCmd.PersistentFlags().StringVarP(&deleteApp, "app", "l", "", "安装指定的应用")
}
