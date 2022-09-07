package input

import (
	"context"
	"customer_partner_match/model"
	"customer_partner_match/pkg/pkg_error"
)

type FloorUseCase interface {
	CreatePartner(ctx context.Context, inputDTO model.NewFloorPartnerDTO) (model.NewFloorPartnerDTO, pkg_error.AppError)
	FindPartners(ctx context.Context, inputDTO model.FloorRequestDTO) ([]model.FloorPartnerResponseDTO, pkg_error.AppError)
}
