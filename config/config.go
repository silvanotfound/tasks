package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err := InitializeSqlite()
	if err != nil {
		return fmt.Errorf("erro ao iniciar o bando de dados")
	}
	logger.Infof("get database", db.Config)
	return nil
}

func GetSqlite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLoger(p)
	return logger
}
