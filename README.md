# badi-events-backend

## Install Dependencies

`go build`

This should use the [Go Modules Tool](https://blog.golang.org/using-go-modules) to automatically download all the necessary dependencies.

## Setup

You will need to copy the `.env.template` file to `.env` and fill in the appropriate values.
See the server `/var/www/html/go/credentials.json` for the credential values.

Once you have done that, when you run the script, you will be prompted to authorize the app and create the appropriate tokens. (Or you can download `token.json` from the server.

All this is to keep credentials out of version control.

## Run Script

`go run *.go`

You need to run all the local `main` package files together with the `run` command.

