package input

import (
	"context"
	"customer_partner_match/pkg/pkgError"
)

type FloorUseCase interface {
	CreatePartner(ctx context.Context)
	FindPartner(ctx context.Context, inputDTO FloorRequestDTO) pkgError.AppError
}

type FloorRequestDTO struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Wood      bool    `json:"wood"`
	Carpet    bool    `json:"carpet"`
	Tiles     bool    `json:"tiles"`
	FloorArea float64 `json:"floor_area"`
	Phone     string  `json:"phone"`
}
