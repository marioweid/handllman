package config

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
    connStr := "user=yourusername dbname=yourdbname password=yourpassword host=db port=5432 sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected to the database")
    return db
}