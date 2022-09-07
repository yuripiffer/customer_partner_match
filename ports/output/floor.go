package output

import (
	"context"
	"customer_partner_match/model"
	"customer_partner_match/pkg/pkg_error"
)

type FloorRepository interface {
	PersistNewPartner(ctx context.Context, inputDTO model.NewFloorPartnerDTO) pkg_error.AppError
	SelectPartners(ctx context.Context, inputDTO model.FloorRequestDTO) ([]model.FloorPartnerResponseDTO, pkg_error.AppError)
}
