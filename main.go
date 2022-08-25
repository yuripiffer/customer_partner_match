package main

import (
	"customer_partner_match/domain/floor"
	"customer_partner_match/infrastructure/config"
	"customer_partner_match/infrastructure/db"
	"customer_partner_match/web"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	config.PopulateEnv()
	dbConnection := db.InitDBConnection(
		config.ENV.DBHost,
		config.ENV.DBPort,
		config.ENV.DBUser,
		config.ENV.DBName,
		config.ENV.DBPassword)
	defer db.CloseDBConnection(dbConnection)

	floorPartnerRepository := db.NewFloorPartnerRepository(dbConnection, "floor_partner_table")

	floorPartnerUseCase := floor.New(floorPartnerRepository)

	r := mux.NewRouter()
	web.ConfigureFloorPartnerRoutes(floorPartnerUseCase, r)
	err := http.ListenAndServe(":85", r)
	if err != nil {
		panic(err)
	}

}
