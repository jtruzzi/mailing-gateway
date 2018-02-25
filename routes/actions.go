package routes

import (
	"net/http"
	"../services"
	"github.com/gorilla/mux"
	"encoding/json"
)


func EmailConfirmation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reservationId := params["reservationId"]

	reservation, err := services.GetReservation(reservationId )
	if err != nil { panic(err) }

	for _, product := range reservation.Products {
		if product.Type != "FLIGHT" { break }

		flightReservation, err2 := services.GetFlightReservation(product.ReservationId)
		if err2 != nil { panic(err); return }

		services.SendEmailConfirmation(reservation, product, flightReservation)
		json.NewEncoder(w).Encode(reservation.Products[0].FlightReservation)
	}
}

