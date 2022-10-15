package main

import (
	"log"
	"net/http"
	"rest/database"
	"rest/routers"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const port = ":8090"

func main() {

	initDB()

	log.Println("Starting http in", port)
	router := mux.NewRouter().StrictSlash(true)
	routers.InitialiseHandlers(router)
	log.Fatal(http.ListenAndServe(port, router))
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "password",
			DB:         "ae_money",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
}
