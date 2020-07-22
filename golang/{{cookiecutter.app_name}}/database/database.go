package database

import (
	"fmt"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/config"
	_ "github.com/lib/pq"
	"database/sql"
	log "github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/log"
)

// InitDB initializes the global database session
func init()  {
	log.Info("Initializing database")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Config().GetString("database.host"),
		config.Config().GetString("database.port"),
		config.Config().GetString("database.user"),
		config.Config().GetString("database.password"),
		config.Config().GetString("database.name"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Info("Successfully connected to database!")
}
