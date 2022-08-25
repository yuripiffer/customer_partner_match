package output

import (
	"context"
	"customer_partner_match/model"
	"customer_partner_match/pkg/pkgError"
)

type FloorRepository interface {
	PersistNewPartner(ctx context.Context, inputDTO model.NewFloorPartnerDTO) pkgError.AppError
	SelectPartners(ctx context.Context, inputDTO model.FloorRequestDTO) (
		[]model.FloorPartnerResponseDTO, pkgError.AppError)
}
