package main

import (
	"context"
	"database/sql"
	"log"
	"stori/internal/adapter/config"
	"stori/internal/adapter/email"
	"stori/internal/adapter/handler/http"
	"stori/internal/adapter/storage/sqlite3/repository"
	"stori/internal/core/service"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open(config.DB.Engine, config.DB.DSN)
	if err != nil {
		log.Fatal(err)
	}

	// Migrations
	if _, err := db.ExecContext(context.Background(), ddl); err != nil {
		log.Fatal(err)
	}

	queries := repository.New(db)

	userRepository := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)

	accountRepository := repository.NewAccountRepository(queries)
	accountService := service.NewAccountService(accountRepository)
	accountHandler := http.NewAccountHandler(accountService)

	emailService := email.NewSesEmail(userRepository, config)

	transactionRepository := repository.NewTransactionRepository(queries)
	transactionService := service.NewTransactionService(transactionRepository, emailService)
	transactionHandler := http.NewTransactionHandler(transactionService)

	router, err := http.NewRouter(
		config,
		*userHandler,
		*accountHandler,
		*transactionHandler,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = router.Serve(config.HTTP.Address)
	if err != nil {
		log.Fatal(err)
	}
}
