package router

import (
	"github.com/autousers/backend_reserve/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/room", middleware.GetAllRoom).Methods("GET", "OPTIONS")

	return router
}
