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
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install命令调用helm安装所有的应用",
	Long:  `install命令调用helm安装所有的应用`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("install called")
		if installAll {
			installAllHelm(root, profile)
		} else if !strings.EqualFold(installSu, "") && strings.EqualFold(installApp, "") {
			installSuHelm(root, profile, installSu)
		} else if !strings.EqualFold(installSu, "") && !strings.EqualFold(installApp, "") {
			installAppHelm(root, profile, installSu, installApp)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	installCmd.PersistentFlags().StringVarP(&charts, "charts", "c", "./", "charts文件的输出目录")
	installCmd.PersistentFlags().StringVarP(&profile, "profile", "p", "profile.yaml", "profile文件的位置")
	installCmd.PersistentFlags().BoolVarP(&installAll, "all", "a", false, "安装所有的charts文件")
	installCmd.PersistentFlags().StringVarP(&installSu, "su", "s", "", "安装指定的服务单元")
	installCmd.PersistentFlags().StringVarP(&installApp, "app", "l", "", "安装指定的应用")
}
