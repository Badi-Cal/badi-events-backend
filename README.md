# badi-events-backend

## Install Dependencies

`go build`

This should use the [Go Modules Tool](https://blog.golang.org/using-go-modules) to automatically download all the necessary dependencies.

## Setup

### Environment Variables

You will need to copy the `.env.template` file to `.env` and fill in the appropriate values.
See the server `/var/www/html/go/credentials.json` for the credential values.

Once you have done that, when you run the script, you will be prompted to authorize the app and create the appropriate tokens. (Or you can download `token.json` from the server.

All this is to keep credentials out of version control.

### Database

Once the environment variables have been set (particularly `BEB_ENV`), 
you will need to initialize the database with a privileged database user (probably `root`):
```shell
setup/setup.sql.sh | sudo mysql -uroot -p
```

The `BEB_ENV` specifies a postfix for the database name depending on the environment (development, test, production, etc).

### Database Tools

We use [Gorm](https://gorm.io/docs/) for ORM and other database queries.

### Database Migrations

Schema Migrations are held in the `migrations` folder. You will need to run them once upon every update.

```shell
go run migrations/*.go
```

You can create a new migration with:
```shell
go run setup/setup.go create_migration [Name your migration]
```

This action will create a new file in the `migrations` folder. It will have an object with two actions `up` and `down`.
The `up` action will be where you will put [migration code](https://gorm.io/docs/migration.html).
The `down` action undoes your migration, which is useful for testing the migration code and undoing changes from a bad deploy.

## Run Script

`go run *.go`

You need to run all the local `main` package files together with the `run` command.

## Requirements

### BadiDate

[BadiDate by Jan Greis](https://github.com/janrg/badiDate/)

### Luxon

Required by BadiDate. [A subproject of Moment.js](https://moment.github.io/luxon/)

## References

- https://gorm.io/docs/
- https://github.com/go-gorm/mysql
- https://developers.google.com/calendar/api/v3/reference
