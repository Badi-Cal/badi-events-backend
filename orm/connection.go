package orm

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *gorm.DB
var databaseName = "notifications"
var Charset = "utf8mb4"
var Collation = "utf8mb4_bin"

type DatabaseConfig struct {
	Username     string
	Password     string
	DatabaseName string
}

func MakeDsn(config DatabaseConfig) string {
	var port int
	var err error
	port, err = strconv.Atoi(os.Getenv("BEB_DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid port. %v\n", err)
	}
	var host = os.Getenv("BEB_DB_HOST")
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		host,
		port,
		config.DatabaseName,
		Charset,
	)
}

func makeBasicDsn() string {
	var username = os.Getenv("BEB_DB_USERNAME")
	var password = os.Getenv("BEB_DB_PASSWORD")
	config := DatabaseConfig{
		Username:     username,
		Password:     password,
		DatabaseName: DatabaseName(),
	}
	return MakeDsn(config)
}

func DatabaseName() string {
	var env = os.Getenv("BEB_ENV")
	return databaseName + "-" + env
}

func Connection() *gorm.DB {
	if connection == nil {
		var err error = nil
		dsn := makeBasicDsn()
		connection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Unable to connect to %s. %v\n", DatabaseName(), err)
		}
	}

	return connection
}
