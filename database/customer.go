package dataservice

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"main.go/model"
)

// decode
func CreateCustomer(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var customer model.Customer
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&customer); err != nil {

		return err
	}
	prod := "select quantity from product where product =?"
	var total int
	err1 := db.QueryRow(prod, customer.Product).Scan(&total)
	if err1 != nil {
		return err1
	}
	if customer.Quantity <= 0 || total < customer.Quantity {
		query := "INSERT INTo customer (Name,Phn,Email,Quantity,Product) VALUES (?,?,?,?)"
		_, err := db.Exec(query, customer.Name, customer.Phn, customer.Email, customer.Quantity, customer.Product)
		if err != nil {
			return err
		}
		//update
		newquantity := total - customer.Quantity
		quantUPDATE := "SET quantity=? where product=?"
		_, err2 := db.Exec(quantUPDATE, newquantity, customer.Product)
		if err2 != nil {
			return err2
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
	return nil
}
