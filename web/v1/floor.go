package v1

import (
	"customer_partner_match/ports/input"
	"net/http"
)

type FloorV1Handler struct {
	UseCase input.FloorUseCase
}

func (h *FloorV1Handler) NewPartner(w http.ResponseWriter, r *http.Request) {}

func (h *FloorV1Handler) FindPartners(w http.ResponseWriter, r *http.Request) {}
