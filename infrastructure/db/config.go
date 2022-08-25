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
