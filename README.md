# Challenge

## Web framework
[gin](https://gin-gonic.com/)

## SQL query builder
[sqlc](https://sqlc.dev/)

## SQL Engine
[sqlite](https://www.sqlite.org/)

## Configure environment
`$ cp .env{.example,} # Then change the values in the .env file`

## Build project
`$ make docker-build`

## Docker Compose
`$ docker compose up`

## General aspects
This project follows the hexagonal architecture and implements the send email
through AWS SES service. For development porpuses uses [air](https://github.com/cosmtrek/air)

## Development
`$ air # After configuring .env file`

## Basic test of the project
```sh
$ cp .env{.example,} # And configure the env variables

$ make docker-build && docker compose up

$ # Call for create a new user in the system
$ curl --location 'http://localhost:8080/v1/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "valid@email.com"
}'

$ # Call for create a new account for the user already created in the last step
$ curl --location 'http://localhost:8080/v1/accounts' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": "valid uuid from the last call", # Change this
    "bank": "Citibanamex",
    "number": "1234567890"
}'

$ # Call for upload the test file to the account created in the last step
$ curl --location 'http://localhost:8080/v1/transactions/upload-file' \
--form 'account_id="valid uuid from last call"' \ # Change this
--form 'file=@"$(pwd)/stori/test.csv"' # Change this for a valid path in the test environment

$ # Call for send summary balance to the user with the account holding the transactions
$ curl --location 'http://localhost:8080/v1/transactions/send-email-summary' \
--header 'Content-Type: application/json' \
--data '{
    "account_id": "valid uuid from last call"
}'
```
