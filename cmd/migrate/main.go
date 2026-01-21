package main

import (
	"github.com/Prakharpandey007/paypocket/config"
	"github.com/Prakharpandey007/paypocket/internal/db"
)

func main() {
	cfg := config.Load()
	db.Connect(cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)
	db.Migrate()
}

