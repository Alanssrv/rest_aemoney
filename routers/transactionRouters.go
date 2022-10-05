package routers

import (
	"rest/controllers"

	"github.com/gorilla/mux"
)

func InitialiseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreateTransaction).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllTransaction).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetTransactionByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdateTransactionByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletTransactionByID).Methods("DELETE")
}
