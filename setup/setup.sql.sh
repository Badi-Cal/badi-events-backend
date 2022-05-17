#!/bin/sh

dir=`dirname $(realpath $0)`

source $dir/../.env

if [ -z $BEB_ENV ]; then
    echo "No environment set in .env"
    exit 1
fi


backtick='`'
database="${backtick}notifications-$BEB_ENV$backtick"
echo "
CREATE DATABASE IF NOT EXISTS $database CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
USE $database;
CREATE TABLE IF NOT EXISTS schema_migrations (
    migration varchar(14) NOT NULL DEFAULT '',
    PRIMARY KEY(migration)
) ENGINE=InnoDB;
CREATE USER IF NOT EXISTS '$BEB_DB_USERNAME'@'localhost' IDENTIFIED BY '$BEB_DB_PASSWORD';
GRANT ALL PRIVILEGES ON $database.* TO '$BEB_DB_USERNAME'@'localhost';
"