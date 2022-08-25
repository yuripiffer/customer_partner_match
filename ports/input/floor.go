package input

import (
	"context"
	"customer_partner_match/model"
	"customer_partner_match/pkg/pkgError"
)

type FloorUseCase interface {
	CreatePartner(ctx context.Context)
	FindPartners(ctx context.Context, inputDTO model.FloorRequestDTO) (
		[]model.FloorPartnerDTO, pkgError.AppError)
}
