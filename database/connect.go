package database

import (
	"fmt"
	"os"

	"sharingvisionbe/ent"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *ent.Client {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return client
}
