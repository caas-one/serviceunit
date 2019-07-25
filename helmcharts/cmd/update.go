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

// updateCmd represents the clean command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update命令用于更新指定服务单元和应用名称的应用",
	Long:  `update命令用于更新指定服务单元和应用名称的应用`,
	Run: func(cmd *cobra.Command, args []string) {
		if updateAll {
			updateAllHelm(root, profile)
		} else if !strings.EqualFold(updateSu, "") && strings.EqualFold(updateApp, "") {
			log.Infof("su=%s, app=%s", updateSu, updateApp)
			updateSuHelm(root, profile, updateSu)
		} else if !strings.EqualFold(updateSu, "") && !strings.EqualFold(updateApp, "") {
			updateAppHelm(root, profile, updateSu, updateApp)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	updateCmd.PersistentFlags().BoolVarP(&updateAll, "all", "a", false, "更新所有的charts文件")
	updateCmd.PersistentFlags().StringVarP(&updateSu, "su", "s", "", "更新指定的服务单元")
	updateCmd.PersistentFlags().StringVarP(&updateApp, "app", "l", "", "更新指定的应用")
	updateCmd.PersistentFlags().StringVarP(&root, "output", "o", "output", "输出charts的文件夹")
	updateCmd.PersistentFlags().StringVarP(&profile, "profile", "p", "profile.yaml", "profile文件的路径")
}
