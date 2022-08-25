package main

import (
	"customer_partner_match/domain/floor"
	"customer_partner_match/infrastructure/config"
	"customer_partner_match/infrastructure/db"
	"fmt"
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

	floorRepository := floor.New(dbConnection, "floor_partner")
	fmt.Println(floorRepository) //remove

}
