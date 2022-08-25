package db

import (
	"context"
	"customer_partner_match/model"
	"customer_partner_match/pkg/pkgError"
	"customer_partner_match/ports/output"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type Repository struct {
	output.FloorRepository
	dbConn *pgx.Conn
	table  string
}

func NewFloorPartnerRepository(dbConn *pgx.Conn, table string) *Repository {
	return &Repository{
		dbConn: dbConn,
		table:  table,
	}
}

func (r *Repository) SelectPartners(ctx context.Context, inputDTO model.FloorRequestDTO) (
	[]model.FloorPartnerDTO, pkgError.AppError) {

	query := buildSelectPartnersQuery(inputDTO)
	rows, err := r.dbConn.Query(ctx, query)
	if err != nil {
		return nil, pkgError.NewServerError("sql query error", err)
	}
	defer rows.Close()

	var floorPartners []model.FloorPartnerDTO
	for rows.Next() {
		var floorPartner model.FloorPartnerDTO
		err = rows.Scan(&floorPartner)
		if err != nil {
			return nil, pkgError.NewServerError("row scan error", err)
		}
		floorPartners = append(floorPartners, floorPartner)
	}
	return floorPartners, nil
}

func buildSelectPartnersQuery(inputDTO model.FloorRequestDTO) string {
	distanceComparison := fmt.Sprintf(`ST_DistanceSphere(
		geometry(POINT(%v,%v)), geometry(POINT(latitude, longitude))
	)`, inputDTO.Latitude, inputDTO.Longitude)

	materialConditions := ""
	if inputDTO.Carpet {
		materialConditions += "AND carpet = TRUE "
	}
	if inputDTO.Wood {
		materialConditions += "AND wood = TRUE "
	}
	if inputDTO.Carpet {
		materialConditions += "AND carpet = TRUE"
	}

	totalPartners := 15 //default
	if inputDTO.TotalPartners > 0 {
		totalPartners = inputDTO.TotalPartners
	}

	query := fmt.Sprintf(`
	SELECT
		id,
		partner,
		rating,
		latitude,
		longitude,
		(select %s) as distance
	FROM
  		pointsTable
	WHERE
	%s <=  operating_radius 
	%s
	ORDER BY
		rating DESC, distance ASC
	LIMIT %v;
	`, distanceComparison, distanceComparison, materialConditions, totalPartners)
	return query
}
