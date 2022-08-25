package floor

import (
	"context"
	"customer_partner_match/infrastructure/db"
	"customer_partner_match/model"
	"customer_partner_match/pkg"
	"customer_partner_match/pkg/pkgError"
	"customer_partner_match/ports/input"
	"errors"
)

type service struct {
	input.FloorUseCase
	floorRepository *db.Repository
}

func New(floorRepository *db.Repository) *service {
	return &service{
		floorRepository: floorRepository,
	}
}
func (s *service) CreatePartner(ctx context.Context) {}

func (s *service) FindPartners(ctx context.Context, inputDTO model.FloorRequestDTO) (
	[]model.FloorPartnerDTO, pkgError.AppError) {

	if !pkg.ValidatePhoneNumber(inputDTO.Phone) {
		return nil, pkgError.NewInputError("missing/invalid field(s)", errors.New("invalid phone number regex"))
	}

	floorPartners, appError := s.floorRepository.SelectPartners(ctx, inputDTO)
	if appError != nil {
		return nil, appError
	}

	return floorPartners, nil
}
