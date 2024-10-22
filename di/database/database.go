package database

import (
	"orcamento/config"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	err      error
)

func New(cfg *config.Config) (*gorm.DB, error) {
	if database != nil {
		return database, nil
	}

	database, err = gorm.Open(mysql.Open(cfg.DatabaseURL), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return database, nil
}
