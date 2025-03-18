# Hangover SMS

Hangover SMS is school Management System.

## Prerequisites

* Install golang latest version from [here](https://go.dev/doc/install)
* Install postgresql latest version 

## How to run?

1. cd preferred_folder and run `git clone git@github.com:robertantonyjaikumar/hangover-sms.git`

2. create `db.env` file and configure this values

````
db_username=username
db_password=password
````

3. create `config.yaml` file and configure this values

````
server:
  port: 8081
database:
  driver: postgres
  host: localhost
  port: 5432
  dbname: dbname
env: "local"
````
4. Run `go mod tidy` command
5. Run `go run main.go`
