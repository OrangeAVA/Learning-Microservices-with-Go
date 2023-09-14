package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabseConnection(conf DBConfig) (*sql.DB, error) {
	dsn := conf.User + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port + ")/" + conf.Database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Error connecting to database")
		return nil, err
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("Error connecting to database")
		return nil, err
	}

	log.Println("database connected")

	return db, nil
}
