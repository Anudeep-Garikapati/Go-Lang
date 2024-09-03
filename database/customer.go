package dataservice

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"main.go/model"
)

func CreateCustomer(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var customer model.Customer
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&customer); err != nil {

		return err
	}
	prod := "select quan from product where product =?"
	var total int
	a := customer.Product
	err1 := db.QueryRow(prod, a).Scan(&total)
	if err1 != nil {
		return err1
	}
	if total >= customer.Quantity && customer.Quantity > 0 {
		query := "INSERT INTo customer (Name,Phn,Email,Quantity,Product) VALUES (?,?,?,?)"
		_, err := db.Exec(query, customer.Name, customer.Phn, customer.Email, customer.Quantity, customer.Product)
		if err != nil {
			return err
		}
		//update
		b := total - customer.Quantity
		quantUPDATE := "SET quantity=? where product=?"
		_, err2 := db.Exec(quantUPDATE, b, customer.Product)
		if err2 != nil {
			return err2
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
	return nil
}
