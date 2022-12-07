package orm

import (
	"database/sql"
	"fmt"
	"log"
)

type databaseRecord struct {
	Database string
}

func DoesDatabaseExist() bool {
	var err error
	dsn := makeBasicDsn()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
	}

	command := fmt.Sprintf(
		"SHOW DATABASES LIKE '%s'",
		DatabaseName(),
	)
	var row string
	err = db.QueryRow(command).Scan(&row)
	if err != nil {
		return false
	}

	return row == DatabaseName()
}
