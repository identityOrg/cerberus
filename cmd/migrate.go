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
	"bufio"
	"fmt"
	"github.com/identityOrg/cerberus/impl/store"
	config2 "github.com/identityOrg/cerberus/setup/config"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/square/go-jose.v2"
	"log"
	"os"
	"strings"
	"time"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runMigration,
}

var (
	force bool
	demo  bool
)

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().BoolVarP(&force, "force", "f", false, "Clean DB tables before create")
	migrateCmd.Flags().BoolVar(&demo, "demo", false, "Create demo client and user")
}

func runMigration(*cobra.Command, []string) {
	if force {
		log.Print("force argument is set, deleting the tables first")
		fmt.Printf("Do you want to continue (Y/n): ")
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		if err != nil {
			log.Println(err)
			log.Fatal("Aborting the migration")
			return
		} else {
			switch char {
			case 'Y':
			case 'y':
				log.Println("continuing the migration with drop table")
			}
		}
	}
	config := config2.NewDBConfig()
	_ = viper.UnmarshalKey("store", &config)
	if ormDB, err := gorm.Open(config.Driver, config.DSN); err != nil {
		log.Fatalf("error opening GORM of type %s and DSN %s", config.Driver, config.DSN)
		return
	} else {
		defer ormDB.Close()

		clientModel := store.ClientDBModel{}
		userModel := store.UserDBModel{}
		keyModel := store.KeyDBModel{}
		tokenModel := store.TokenDBModel{}
		sessionModel := sessionTable{}
		if force {
			handleError(ormDB.DropTableIfExists(clientModel).Error, "drop", clientModel.TableName())
			handleError(ormDB.DropTableIfExists(userModel).Error, "drop", userModel.TableName())
			handleError(ormDB.DropTableIfExists(keyModel).Error, "drop", keyModel.TableName())
			handleError(ormDB.DropTableIfExists(tokenModel).Error, "drop", tokenModel.TableName())
			handleError(ormDB.DropTableIfExists(sessionModel).Error, "drop", "sessions")
		}

		handleError(ormDB.AutoMigrate(clientModel).Error, "migrate", clientModel.TableName())
		handleError(ormDB.AutoMigrate(userModel).Error, "migrate", userModel.TableName())
		handleError(ormDB.AutoMigrate(keyModel).Error, "migrate", keyModel.TableName())
		handleError(ormDB.AutoMigrate(tokenModel).Error, "migrate", tokenModel.TableName())
		handleError(ormDB.Table("sessions").AutoMigrate(sessionModel).Error, "migrate", "sessions")

		if demo {
			log.Println("Creating demo client with client_id=client and client_secret=client")
			client := store.NewClientDBModel()

			client.ClientID = "client"
			client.ClientSecret = "client"
			client.Public = false
			client.SetApprovedGrantTypes(strings.Split("authorization_code|password|refresh_token|client_credentials|implicit", "|"))
			client.SetRedirectURIs([]string{"http://localhost:8080/redirect"})
			client.SetApprovedScopes(strings.Split("openid|offline|offline_access", "|"))
			client.SetIDTokenSigningAlg(jose.RS256)

			err := ormDB.Create(client).Error
			if err != nil {
				log.Fatalf("failed to create demo client")
			} else {
				log.Println("demo client created")

				log.Println("Creating demo user with username=user and password=user")
				password, _ := bcrypt.GenerateFromPassword([]byte("user"), 12)
				user := store.UserDBModel{
					Username:          "user",
					Password:          string(password),
					Locked:            false,
					Blocked:           false,
					WrongAttemptCount: 0,
				}
				err = ormDB.Create(&user).Error
				if err != nil {
					log.Fatalf("failed to create demo user")
				} else {
					log.Println("demo user created")
				}
			}
		}
		log.Println("Migration operation complete")
	}
}

func handleError(err error, op string, name string) {
	if err != nil {
		log.Fatalf("failed to migrate %s %s", op, name)
	} else {
		log.Printf("Auto migrated table %s %s", op, name)
	}
}

type sessionTable struct {
	Id        string    `gorm:"primary_key"`
	Data      string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:""`
	UpdatedAt time.Time `gorm:""`
	ExpiresAt time.Time `gorm:"index"`
}
