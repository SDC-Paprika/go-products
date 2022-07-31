package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

var db DB

func Init() {
	dbConfig := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	var err error
	db, err = ConnectDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB(connString string) (DB, error) {
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return DB{}, err
	}

	if err = db.Ping(context.Background()); err != nil {
		return DB{}, err
	}

	return DB{db}, nil
}

func GetDB() DB {
	return db
}
