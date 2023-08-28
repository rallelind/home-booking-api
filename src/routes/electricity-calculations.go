package routes

import (
	"home-booking-api/src/controllers"

	"github.com/gorilla/mux"
)

func RegisterElectricityRoutes(r *mux.Router) {
	r.Handle("/electricity/prices", controllers.GetElectricityPrices()).Methods("GET")
}
