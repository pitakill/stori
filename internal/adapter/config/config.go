package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	App struct {
		Name string
		Env  string
	}

	DB struct {
		Engine string
		DSN    string
	}

	HTTP struct {
		Address        string
		AllowedOrigins string
	}

	SES struct {
		TemplateName string
		EmailFrom    string
	}

	AWS struct {
		SES *SES
	}

	Container struct {
		App  *App
		DB   *DB
		HTTP *HTTP
		AWS  *AWS
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	db := &DB{
		Engine: os.Getenv("DB_ENGINE"),
		DSN:    os.Getenv("DB_DSN"),
	}

	http := &HTTP{
		Address:        os.Getenv("HTTP_ADDRESS"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}

	ses := &SES{
		TemplateName: os.Getenv("AWS_SES_TEMPLATE_NAME"),
		EmailFrom:    os.Getenv("AWS_SES_EMAIL_FROM"),
	}

	aws := &AWS{
		SES: ses,
	}

	return &Container{
		app,
		db,
		http,
		aws,
	}, nil
}
