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
