package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

func InitDBConnection(host string, port string, user string, name string, password string) *pgx.Conn {
	connConfig, err := pgx.ParseConfig(getDBConnectionString(host, port, user, name, password))
	if err != nil {
		panic(err)
	}

	connection, err := pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		panic(err)
	}
	return connection
}

func CloseDBConnection(dbConnection *pgx.Conn) {
	if dbConnection != nil {
		dbConnection.Close(context.Background())
	}
}

func getDBConnectionString(host string, port string, user string, name string, password string) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		host, port, user, name, password)
}

//Only for documentation
func createFloorPartnerTableQuery() {
	_ = fmt.Sprint(`
	CREATE TABLE floor_partner_table (
		id UUID primary key, 
		partner VARCHAR(255),
		rating numeric not null check (rating between 1 and 5), 
		operating_radius numeric CHECK (operating_radius > 0), 
		latitude numeric not null check (latitude between -180 and 180), 
		longitude numeric not null check (longitude between -180 and 180), 
		wood bool, 
		carpet bool, 
		tiles bool);`)
}
