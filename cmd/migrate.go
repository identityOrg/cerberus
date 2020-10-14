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
	"github.com/identityOrg/cerberus-core"
	"github.com/identityOrg/cerberus/setup"
	config2 "github.com/identityOrg/cerberus/setup/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Long: `This command will migrate the DB to its desired state:

Do not rely on it fully, its is suitable to only create
the DB for the first time. Consecutive call may give
un-predictable result.`,
	Run: runMigration,
}

var (
	force bool
	reset bool
	demo  bool
)

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().BoolVarP(&force, "force", "f", false, "force non-interactive db reset")
	migrateCmd.Flags().BoolVarP(&reset, "reset", "r", false, "reset all db tabled")
	migrateCmd.Flags().BoolVar(&demo, "demo", false, "create demo client and user")
}

func runMigration(_ *cobra.Command, arg []string) {
	config := config2.NewDBConfig()
	_ = viper.UnmarshalKey("store", &config)
	if ormDB, err := setup.NewGormDB(config); err != nil {
		log.Fatalf("error opening GORM of type %s and DSN %s", config.Driver, config.DSN)
		return
	} else {
		if debug {
			ormDB = ormDB.Debug()
		}
		err := core.SetupDBStructure(ormDB, reset, force)
		if err != nil {
			log.Fatal("migration failed", err)
			return
		}
		//sessionModel := sessionTable{}
		//if force {
		//	handleError(ormDB.DropTableIfExists(sessionModel).Error, "drop", "sessions")
		//}
		//handleError(ormDB.Table("sessions").AutoMigrate(sessionModel).Error, "migrate", "sessions")

		if demo {
			coreConfig := config2.NewCoreConfig()
			sdkConfig := config2.NewSDKConfig()
			redirectUri := ""
			if len(arg) > 0 {
				redirectUri = arg[0]
			}
			err := core.SetupDemoData(ormDB, coreConfig, sdkConfig, redirectUri)
			if err != nil {
				log.Fatal("migration failed", err)
				return
			}
		}
		log.Println("migration operation complete")
	}
}

//func handleError(err error, op string, name string) {
//	if err != nil {
//		log.Fatalf("failed to migrate %s %s", op, name)
//	} else {
//		log.Printf("Auto migrated table %s %s", op, name)
//	}
//}
//
//type sessionTable struct {
//	Id        string    `gorm:"primary_key"`
//	Data      string    `gorm:"type:text"`
//	CreatedAt time.Time `gorm:""`
//	UpdatedAt time.Time `gorm:""`
//	ExpiresAt time.Time `gorm:"index"`
//}
