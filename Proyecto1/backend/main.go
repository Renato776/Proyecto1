package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"test_servers/backend/db"
	"test_servers/backend/handles"
)

func main() {
	if _, err := db.GetConnection(); err != nil {
		log.Fatal("Error to connect a DB")
	}
	router := mux.NewRouter()
	router.HandleFunc("/reports", handles.CreateNewReport).Methods(http.MethodPost)
	router.HandleFunc("/reports", handles.GetReports).Methods(http.MethodGet)
	router.HandleFunc("/reports/{id}", handles.GetReports).Methods(http.MethodGet)
	router.HandleFunc("/reports/carnet/{carnet:[0-9]*}", handles.GetReportsByCarnet).Methods(http.MethodGet)
	http.Handle("/", router)
	fmt.Println("Server on port 4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
