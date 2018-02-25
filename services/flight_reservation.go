package services

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"../models"
	"os"
	"fmt"
)

func GetFlightReservation(reservationId string) (models.FlightReservation, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://api.basset.ws/flights/reservations/" + reservationId  + "?site=AR&channel=WEB", nil)

	request.Header.Add("X-Api-Key", os.Getenv("X_API_KEY"))
	request.Header.Add("x-client-id", os.Getenv("X_CLIENT_ID"))

	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The http request failed with error %s/n", err)
		return models.FlightReservation{}, err
	}

	data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var flightReservation models.FlightReservation
	errMarshall := json.Unmarshal([]byte(data), &flightReservation )
	if errMarshall != nil {
		return models.FlightReservation{}, errMarshall
	}
	return flightReservation, nil
}
