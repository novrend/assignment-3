package router

import (
	"assignment-3/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/data", controllers.UpdateData).Methods("PUT")
	return router
}
