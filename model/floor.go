package model

import "github.com/gofrs/uuid"

type FloorRequestDTO struct {
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Wood          bool    `json:"wood"`
	Carpet        bool    `json:"carpet"`
	Tiles         bool    `json:"tiles"`
	FloorArea     float64 `json:"floor_area"`
	Phone         string  `json:"phone"`
	TotalPartners int     `json:"total_partners"`
}

type FloorPartnerDTO struct {
	ID        uuid.UUID `json:"id"`
	Partner   string    `json:"partner"`
	Rating    float64   `json:"rating"`
	Distance  float64   `json:"distance"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
}
