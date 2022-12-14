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

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	sql_cmd := "InsertTransaction(?, ?, ?, ?)"

	stmt, err := database.Connector.Prepare("CALL " + sql_cmd)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	var transaction entity.Transaction
	json.Unmarshal(body, &transaction)

	_, err = stmt.Exec(transaction.Name, transaction.Value, transaction.Type, transaction.Category)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New post was created")
}

func GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var transactions []entity.Transaction

	sql_cmd := "GetAllTransactions"

	result, err := database.Connector.Query("CALL " + sql_cmd)
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

	var transaction = InternalGetTransactionByID(params["id"])

	if (transaction != entity.Transaction{}) {
		json.NewEncoder(w).Encode(transaction)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
	}

}

func UpdateTransactionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var baseTransaction = InternalGetTransactionByID(params["id"])

	if (baseTransaction != entity.Transaction{}) {
		sql_cmd := "UpdateTransaction(?, ?, ?, ?, ?)"
		stmt, err := database.Connector.Prepare("CALL " + sql_cmd)
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
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
	}

}

func DeleteTransactionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var transaction = InternalGetTransactionByID(params["id"])

	if (transaction != entity.Transaction{}) {
		sql_cmd := "DeleteTransactionByID(?)"
		stmt, err := database.Connector.Prepare("CALL " + sql_cmd)
		if err != nil {
			panic(err.Error())
		}

		_, err = stmt.Exec(params["id"])
		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
	}
}

func InternalGetTransactionByID(id string) entity.Transaction {
	sql_cmd := "GetTransactionById(?)"

	result, err := database.Connector.Query("CALL "+sql_cmd, id)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var transaction entity.Transaction
	defer result.Close()

	_ = scan.Row(&transaction, result)

	return transaction
}
