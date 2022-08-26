package model

type FloorRequestDTO struct {
	Latitude      float64 `schema:"latitude"`
	Longitude     float64 `schema:"longitude"`
	FloorArea     float64 `schema:"floor_area"`
	Phone         string  `schema:"phone"`
	TotalPartners int     `schema:"total_partners"`
	Wood          bool    `schema:"wood"`
	Carpet        bool    `schema:"carpet"`
	Tiles         bool    `schema:"tiles"`
}

type FloorPartnerResponseDTO struct {
	ID        string  `json:"id"`
	Partner   string  `json:"partner"`
	Rating    float64 `json:"rating"`
	Distance  float64 `json:"distance"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type NewFloorPartnerDTO struct {
	ID              string  `json:"id"`
	Partner         string  `json:"partner"`
	Rating          float64 `json:"rating"`
	OperatingRadius int     `json:"operating_radius"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Wood            bool    `json:"wood,omitempty"`
	Carpet          bool    `json:"carpet,omitempty"`
	Tiles           bool    `json:"tiles,omitempty"`
}
