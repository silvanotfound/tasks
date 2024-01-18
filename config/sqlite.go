package config

import (
	"os"
	"silvanotfound/tasks/schemas"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := GetLogger("SQLite")
	dbPath := "./db/tasks.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Create data file ")
		err = os.Mkdir("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Erro ao iniciar o bando de dados", err.Error())
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Task{})
	if err != nil {
		logger.Errorf("Automigration error", err)
	}
	return db, nil
}
