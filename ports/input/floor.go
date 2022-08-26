package input

import (
	"context"
	"customer_partner_match/model"
	"customer_partner_match/pkg/pkgError"
)

type FloorUseCase interface {
	CreatePartner(ctx context.Context, inputDTO model.NewFloorPartnerDTO) (model.NewFloorPartnerDTO, pkgError.AppError)
	FindPartners(ctx context.Context, inputDTO model.FloorRequestDTO) ([]model.FloorPartnerResponseDTO, pkgError.AppError)
}
