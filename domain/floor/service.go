package floor

import (
	"context"
	"customer_partner_match/pkg/pkgError"
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
func (s *service) CreatePartner(ctx context.Context) {}

func (s *service) FindPartner(ctx context.Context, inputDTO input.FloorRequestDTO) pkgError.AppError {

	return nil
}
