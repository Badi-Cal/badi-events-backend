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
# Run all un-run migrations
go run migrations/*.go migrate
```

You can create a new migration with:
```shell
go run setup/setup.go create_migration [Name your migration]
```

This action will create a new file in the `migrations` folder and the migration will be given a Version number or ID.

You can list all existing migration Versions with:

```shell
go run migrations/*.go list
```
It will have an object with two actions `up` and `down`.
The `up` action will be where you will put [migration code](https://gorm.io/docs/migration.html).
```shell
go run migrations/*.go migrate:up VERSION_NUMBER
```
The `down` action undoes your migration, which is useful for testing the migration code and undoing changes from a bad deploy.
```shell
go run migrations/*.go migrate:down VERSION_NUMBER
```

## Run Script

`go run *.go`

You need to run all the local `main` package files together with the `run` command.

## Testing

Run current tests:

```
go test ./controllers/
```

The controller tests stub the model layer of the code using [mockgen](https://github.com/golang/mock), rather than using database seeds (at the moment).

To create or update the stubs when one has updated the model layer, run (inserting the correct model file):
```
mockgen -source=models/notifications.go -destination=mock_models/notifications.go
```

`mockgen` uses interfaces to create the mock, so make sure you have updated the interfaces in the model layer if you add any functions.

## Requirements

### BadiDate

[BadiDate by Jan Greis](https://github.com/janrg/badiDate/)

### Luxon

Required by BadiDate. [A subproject of Moment.js](https://moment.github.io/luxon/)

## References

- https://gorm.io/docs/
- https://github.com/go-gorm/mysql
- https://developers.google.com/calendar/api/v3/reference
