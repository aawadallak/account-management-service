package database

import (
	"context"
	"fmt"
	"latest/config"
	"log"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

type Database struct {
	databaseUser     string
	databasePassword string
	databaseSSL      string
	databaseName     string
	databaseHost     string
}

func createDatabase() Database {
	return Database{
		databaseUser:     config.GetConfig().DatabaseUser,
		databasePassword: config.GetConfig().DatabasePass,
		databaseSSL:      config.GetConfig().DatabaseSSL,
		databaseName:     config.GetConfig().DatabaseName,
		databaseHost:     config.GetConfig().DatabaseHost,
	}

}

var database *pgx.Conn

func Init() {

	_db := createDatabase()

	// urlExample := "postgres://username:password@localhost:5432/database_name"

	str := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", _db.databaseUser, _db.databasePassword, _db.databaseHost, _db.databaseName)

	db, err := pgx.Connect(context.Background(), str)

	if err != nil {
		log.Fatalln(err)
	}

	database = db

}

func GetDatabase() *pgx.Conn {
	return database
}
