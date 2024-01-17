package main

import (
	"badi-cal/badi-events-backend/orm"
	"badi-cal/badi-events-backend/terminal"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type SchemaMigration struct {
	Migration string `gorm:"primaryKey"`
}

type MigrationAction func(db *gorm.DB)
type Migration struct {
	name    string
	version string
	up      MigrationAction
	down    MigrationAction
}

// var argLength int = -1
var versions []Migration

func (migration *Migration) Migrate() {
	db := orm.Connection()
	var schema_migration SchemaMigration
	if result := db.Where("migration = ?", migration.version).Take(&schema_migration); result.Error != nil {
		fmt.Printf("Running migration '%s'", migration.version)
		migration.up(db)
		schema_migration = SchemaMigration{Migration: migration.version}
		result := db.Create(&schema_migration)
		if result.Error != nil {
			log.Fatalf("Unable to mark migration. %v\n", result.Error)
		}
	} else {
		fmt.Printf("Skipping migration '%s'; already run.\n", migration.version)
		fmt.Printf("Error %v\n", result.Error)
	}
}

func (migration *Migration) Rollback() {
	db := orm.Connection()
	var schema_migration SchemaMigration
	if result := db.Where("migration = ?", migration.version).Take(&schema_migration); result.Error != nil {
		fmt.Printf("Unable to rollback migration '%s'; has not been run.\n", migration.version)
	} else {
		fmt.Printf("Rolling back migration '%s'\n", migration.version)
		migration.down(db)
		db.Where("migration = ?", migration.version).Delete(&schema_migration)
	}
}

func registerMigration(migration Migration) {
	versions = append(versions, migration)
}

func main() {
	if terminal.ArgCount() == 0 {
		fmt.Println("No commands given")
		return
	}

	command := terminal.GetArg(0)
	if command == "list" {
		for _, m := range versions {
			fmt.Printf("%s - %s\n", m.version, m.name)
		}
	} else if command == "migrate" {
		fmt.Printf("Arrived %v\n", versions)
		for _, migration := range versions {
			migration.Migrate()
		}
	} else if command == "migrate:up" {
		version := terminal.GetArg(1)
		if version == "" {
			fmt.Println("Version required, not given.")
			return
		}
		var migration Migration
		for _, m := range versions {
			if m.version == version {
				migration = m
				break
			}
		}
		migration.Migrate()
	} else if command == "migrate:down" {
		version := terminal.GetArg(1)
		if version == "" {
			fmt.Println("Version required, not given.")
			return
		}
		var migration Migration
		for _, m := range versions {
			if m.version == version {
				migration = m
				break
			}
		}
		migration.Rollback()
	} else {
		fmt.Printf("Unknown command '%s'.\n", command)
	}
}
