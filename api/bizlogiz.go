package api

import (
	"database/sql"
	"net/http"

	dataservice "main.go/database"
)

//func CreateBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
//	return dataservice.CreateBook(db, w, r)

// }
func CreateCustomerLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.CreateCustomer(db, w, r)
}
