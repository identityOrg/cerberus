/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// initconfigCmd represents the initconfig command
var initconfigCmd = &cobra.Command{
	Use:   "initconfig",
	Short: "A brief description of your command",
	Long: `Initialize config file:

Merge fill the config file with all available
configurations. It uses default value provided
with the application, along with the values changed
through various config sources.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Error("config file name not specified")
		}
		viper.SetConfigType("toml")
		err := viper.WriteConfigAs(args[0])
		if err != nil {
			log.Error(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initconfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initconfigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initconfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
