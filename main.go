package main

import (
	"github.com/enesusta/tzone/county"
	"github.com/enesusta/tzone/province"
	"github.com/enesusta/tzone/town"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/provinces", province.GetCities).Methods("GET")
	r.HandleFunc("/provinces/{provinceName}", province.GetCity).Methods("GET")
	r.HandleFunc("/counties", county.GetCounties).Methods("GET")
	r.HandleFunc("/counties/{provinceName}", county.GetCounty).Methods("GET")
	r.HandleFunc("/towns", town.GetTowns).Methods("GET")
	r.HandleFunc("/towns/{provinceName}", town.GetTown).Methods("GET")
	r.HandleFunc("/towns/{provinceName}/{countyName}", town.GetSpecificTown).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
