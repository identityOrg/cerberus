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
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cerberus",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is cerberus.yaml at work dir)")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "enable debug message")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".cerberus" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigName("cerberus")
	}
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	setDefaultConfiguration()
}

func setDefaultConfiguration() {
	viper.SetDefault("server", map[string]interface{}{"port": "localhost:8080", "debug": false, "demo": false})
	viper.SetDefault("db", map[string]string{"driver": "sqlite3", "dsn": "test.db"})
	viper.SetDefault("secret.SessionSecret", "jdhfbwjhebajwhevbahwevbahevbajwhevblawev")
	viper.SetDefault("secret.TokenSecret", "jhkawhjebawhebvajebvkjahebvkjahebvkjebvj")
	viper.SetDefault("provider", map[string]interface{}{
		"Issuer":                   "http://localhost:8080",
		"AuthCodeLifespan":         10 * time.Minute,
		"AccessTokenLifespan":      1 * time.Hour,
		"RefreshTokenLifespan":     30 * 24 * time.Hour,
		"AccessTokenEntropy":       20,
		"AuthorizationCodeEntropy": 20,
		"RefreshTokenEntropy":      20,
		"StateParamMinimumEntropy": 10,
		"GlobalConsentRequired":    true,
		"PKCEPlainEnabled":         false,
	})
	viper.SetDefault("core", map[string]interface{}{
		"EncryptionKey":          "dafefascdaewwevawevwevdwfef",
		"MaxInvalidLoginAttempt": 3,
		"InvalidAttemptWindow":   time.Minute * 5,
		"TOTPSecretLength":       6,
	})
}
