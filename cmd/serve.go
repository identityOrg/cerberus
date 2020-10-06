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
	"context"
	"github.com/identityOrg/cerberus/setup"
	"github.com/identityOrg/cerberus/setup/config"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Cerberus server",
	Long: `This command serves the Cerberus server:

Start the server on configured port as per configuration
found in designated file.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = viper.BindPFlags(cmd.Flags())
		e, err := setup.CreateEchoServer()
		if err != nil {
			log.Fatal("failed to setup echo server", err)
			return
		}
		serverConfig := config.NewServerConfig()
		// Start server
		go func() {
			if err := e.Start(serverConfig.Port); err != nil {
				e.Logger.Info("shutting down the server")
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 10 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolVar(&demo, "demo", false, "Create demo client and user")
	serveCmd.Flags().StringP("addr", "a", "", "address to start server on")
}
