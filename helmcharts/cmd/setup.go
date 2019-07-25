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
	"github.com/spf13/cobra"
	"gitlab.yeepay.com/yce/helmcharts/setup"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "setup命令根据提供的namespace生成profile.yaml",
	Long:  `setup命令根据提供的namespace生产profile.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Setup the profile: apiserver=%s, output=%s", setup.APIServer, outputProfile)
		err := setupProfile(setup.APIServer, setup.Ca, setup.Cert, setup.Key, outputProfile, srcNamespace, dstNamespace)
		if err != nil {
			log.Errorf("build command setup error: err=%s", err)
		}
		log.Infof("Generate the profile %s successfully!", outputProfile)
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	setupCmd.PersistentFlags().StringVarP(&outputProfile, "output", "o", "profile.yaml", "profile文件的路径")
	setupCmd.PersistentFlags().StringVarP(&srcNamespace, "src", "s", "cicd-default", "源namespace名称")
	setupCmd.PersistentFlags().StringVarP(&dstNamespace, "dst", "d", "newspace", "目标（新的）namespace名称")
	// buildCmd.PersistentFlags().StringVarP(&APIServer, "apiserver", "a", "10.151.33.87:6443", "Host of CICD ApiServer")
	// buildCmd.PersistentFlags().StringVarP(&Ca, "ca", "c", "ca.crt", "Path to ca.cert")
	// buildCmd.PersistentFlags().StringVarP(&Cert, "cert", "e", "client.crt", "Path to client.crt")
	// buildCmd.PersistentFlags().StringVarP(&Key, "key", "k", "client.key", "Path to client.key")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
