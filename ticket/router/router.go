package router

import (
	"github.com/gorilla/mux"
	"github.com/thanhyarn/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/tickets", controller.CreateTicket).Methods("POST")
	router.HandleFunc("/api/tickets/{id}", controller.UpdateTicket).Methods("PUT")
	router.HandleFunc("/api/tickets/{id}/assign/{employeeID}", controller.AssignTicket).Methods("PUT")
	router.HandleFunc("/api/tickets/{id}/status", controller.UpdateTicketStatus).Methods("PUT")
	router.HandleFunc("/api/tickets", controller.GetTicketsByStatus).Queries("status", "{status}").Methods("GET")
	router.HandleFunc("/api/tickets", controller.GetTicketsByDepartment).Queries("department", "{department}").Methods("GET")
	router.HandleFunc("/api/tickets/{id}", controller.GetTicketDetails).Methods("GET")
	router.HandleFunc("/api/tickets/{id}", controller.DeleteTicket).Methods("DELETE")
	router.HandleFunc("/api/tickets/{id}/subtickets", controller.CreateSubtickets).Methods("POST")

	return router
}
