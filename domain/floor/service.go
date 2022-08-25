package floor

import "github.com/jackc/pgx/v4"

type service struct {
	dbConn *pgx.Conn
	table  string
}

func New(dbConn *pgx.Conn, table string) *service {
	return &service{
		dbConn: dbConn,
		table:  table,
	}
}

func (s *service) FindPartner() {}

func (s *service) CreatePartner() {}
