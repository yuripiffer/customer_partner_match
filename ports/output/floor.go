package output

import (
	"context"
	"customer_partner_match/model"
	"customer_partner_match/pkg/pkgError"
)

type FloorRepository interface {
	SelectPartners(ctx context.Context, inputDTO model.FloorRequestDTO) (
		[]model.FloorPartnerDTO, pkgError.AppError)
}
