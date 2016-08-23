package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type DbInitializer interface {
	InitDB(connection string)
	InitSchema()
}

func (i *PostgreDB) InitDB(connection string) {
	var err error
	i.DB, err = gorm.Open("postgres", connection)
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	i.DB.LogMode(true)
}

