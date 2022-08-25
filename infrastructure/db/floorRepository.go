package db

import (
	"context"
	"customer_partner_match/model"
	"customer_partner_match/pkg/pkgError"
	"customer_partner_match/ports/output"
	"errors"
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
func (r *Repository) PersistNewPartner(ctx context.Context, inputDTO model.NewFloorPartnerDTO) pkgError.AppError {
	query := r.buildPersistPartnerQuery(inputDTO)
	exec, err := r.dbConn.Exec(ctx, query)
	if err != nil {
		return pkgError.NewServerError("postgres insert", err)
	}

	if exec.RowsAffected() != 1 {
		return pkgError.NewServerError("postgres insert",
			errors.New(fmt.Sprintf("rows affected != 1, rows: %v", exec.RowsAffected())))
	}
	return nil
}

func (r *Repository) SelectPartners(ctx context.Context, inputDTO model.FloorRequestDTO) (
	[]model.FloorPartnerResponseDTO, pkgError.AppError) {

	query := r.buildSelectPartnersQuery(inputDTO)
	rows, err := r.dbConn.Query(ctx, query)
	if err != nil {
		return nil, pkgError.NewServerError("sql query error", err)
	}
	defer rows.Close()

	var floorPartners []model.FloorPartnerResponseDTO
	for rows.Next() {
		var floorPartner model.FloorPartnerResponseDTO
		err = rows.Scan(&floorPartner)
		if err != nil {
			return nil, pkgError.NewServerError("row scan error", err)
		}
		floorPartners = append(floorPartners, floorPartner)
	}
	return floorPartners, nil
}

func (r *Repository) buildPersistPartnerQuery(inputDTO model.NewFloorPartnerDTO) string {
	query := fmt.Sprintf(`
	INSERT INTO
		%s
		(id, partner, rating, operating_radius, latitude, longitude, wood, carpet, tiles)
	VALUES
		(%s, %s, %v, %v, %v, %v, %v, %v, %v);
	`, r.table,
		inputDTO.ID,
		inputDTO.Partner,
		inputDTO.Rating,
		inputDTO.OperatingRadius,
		inputDTO.Latitude,
		inputDTO.Longitude,
		inputDTO.Wood,
		inputDTO.Carpet,
		inputDTO.Tiles)

	return query
}

func (r *Repository) buildSelectPartnersQuery(inputDTO model.FloorRequestDTO) string {
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
  		%s
	WHERE
	%s <=  operating_radius 
	%s
	ORDER BY
		rating DESC, distance ASC
	LIMIT %v;
	`, distanceComparison, r.table, distanceComparison, materialConditions, totalPartners)
	return query
}
