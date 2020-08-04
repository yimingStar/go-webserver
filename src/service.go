package main

import (
    _ "fmt"
	_ "net/http"
	"github.com/gorilla/mux"
	logger "go-webservice/src/logger"
	api "go-webservice/src/api"
)

func main() {
	logger.InitLogger("webservice")
	router := mux.NewRouter()
	// Read-all
	router.HandleFunc("/test", api.TestAlive).Methods("GET")

	// Swagger
	// router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	logger.WriteLog("info", "Starting server at port 5566")
	http.ListenAndServe(":5566", router)
}