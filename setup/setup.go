package main

import (
	"badi-cal/badi-events-backend/orm"
	"badi-cal/badi-events-backend/terminal"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func CreateMigration(name string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	currentTime := time.Now()
	timestamp := currentTime.Format("20060102150405")
	filename := fmt.Sprintf(
		"%s_%s.go",
		timestamp,
		name,
	)
	path := filepath.Join(cwd, "migrations", filename)
	file, err := os.Create(path)

	migrationName := "migration" + timestamp + "_" + name

	// // close the file with defer
	defer file.Close()
	// // write a string
	file.WriteString(
		`package main


import (
	"badi-cal/badi-events-backend/orm"

	"gorm.io/gorm"
)

var ` + migrationName + ` = Migration{
	name: "` + migrationName + `",
	version: "` + timestamp + `",
	up: func (db *gorm.DB) {

	},
	down: func (db *gorm.DB) {

	},
}

func init() {
	registerMigration(` + migrationName + `)
}
`)
}

func main() {
	if terminal.ArgCount() == 0 {
		fmt.Println("No command given")
		return
	}

	if !orm.DoesDatabaseExist() {
		fmt.Printf("Database %s does not exist or user does not have privileges.\n", orm.DatabaseName())
		return
	}

	command := terminal.GetArg(0)
	if command == "create_migration" {
		if terminal.ArgCount() < 1 {
			fmt.Println("No name given.")
			return
		}
		name := terminal.GetArg(1)
		CreateMigration(name)
	} else {
		fmt.Printf("Unknown command '%s'.\n", command)
	}
}
