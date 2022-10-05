package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest/database"
	"rest/entity"

	"github.com/blockloop/scan"
	"github.com/gorilla/mux"
)

const table = "Transactions"

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	sql_cmd := fmt.Sprintf("INSERT INTO %s VALUES(?, ?, ?, ?, ?)", table)

	stmt, err := database.Connector.Prepare(sql_cmd)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	var transaction entity.Transaction
	json.Unmarshal(body, &transaction)

	_, err = stmt.Exec(transaction.ID, transaction.Name, transaction.Value, transaction.Type, transaction.Category)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New post was created")
}

func GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var transactions []entity.Transaction

	sql_cmd := fmt.Sprintf("SELECT * FROM %s", table)

	result, err := database.Connector.Query(sql_cmd)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()
	err2 := scan.Rows(&transactions, result)
	if err2 != nil {
		panic(err2.Error())
	}

	json.NewEncoder(w).Encode(transactions)
}

func GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	sql_cmd := fmt.Sprintf("SELECT * FROM %s WHERE ID = ?", table)

	result, err := database.Connector.Query(sql_cmd, params["id"])
	if err != nil {
		panic(err.Error())
	}

	var transaction entity.Transaction
	defer result.Close()
	err = scan.Row(&transaction, result)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(transaction)

}

func UpdateTransactionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	sql_cmd := fmt.Sprintf("UPDATE %s SET Name = ?, Value = ?, Type = ?, Category = ? WHERE ID = ?", table)
	stmt, err := database.Connector.Prepare(sql_cmd)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	var transaction entity.Transaction
	json.Unmarshal(body, &transaction)

	name := transaction.Name
	value := transaction.Value
	typec := transaction.Type
	category := transaction.Category

	_, err = stmt.Exec(name, value, typec, category, params["id"])
	if err != nil {
		panic(err.Error())
	}
}

func DeletTransactionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	sql_cmd := fmt.Sprintf("DELETE FROM %s WHERE ID = ?", table)
	stmt, err := database.Connector.Prepare(sql_cmd)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])
}
