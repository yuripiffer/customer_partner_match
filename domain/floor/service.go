package floor

import (
	"customer_partner_match/ports/input"
	"github.com/jackc/pgx/v4"
)

type service struct {
	input.FloorUseCase
	dbConn *pgx.Conn
	table  string
}

func New(dbConn *pgx.Conn, table string) *service {
	return &service{
		dbConn: dbConn,
		table:  table,
	}
}
func (s *service) CreatePartner() {}

func (s *service) FindPartner() {}
