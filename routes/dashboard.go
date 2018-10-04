package routes

import (
	"github.com/gorilla/mux"
	"../ctx/dashboardCtx"
)

func DashboardRoutes(r *mux.Router) {
    r.HandleFunc("/dashboard-hotelcalav", dashboardCtx.Api)
}