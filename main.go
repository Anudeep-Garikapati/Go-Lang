package main

import (
	"database/sql"
	"log"
	"net/http"

	"main.go/api"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:anudeep@tcp(localhost:3306)/library"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	api.RegisterRouter(db)
	log.Println("server strt on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
