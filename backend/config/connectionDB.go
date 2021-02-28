package config

import (
	"database/sql"
	"fmt"
	"log"
	"testMekar/tools"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() (*sql.DB, error) {
	dbUser := tools.GetEnv("DB_USER", "root")
	dbPass := tools.GetEnv("DB_PASS", "password")
	dbHost := tools.GetEnv("DB_HOST", "localhost")
	dbPort := tools.GetEnv("DB_PORT", "3306")
	dbName := tools.GetEnv("DB_NAME", "schema")

	loadData := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, _ := sql.Open("mysql", loadData)

	err := db.Ping()
	if err != nil {
		log.Print(err)
	} else {
		fmt.Printf("Database %v Connected!!!\n", dbName)
	}

	return db, nil
}
