package web

import (
	"customer_partner_match/ports/input"
	v1 "customer_partner_match/web/v1"
	"github.com/gorilla/mux"
)

func ConfigureFloorPartnerRoutes(useCase input.FloorUseCase, r *mux.Router) {
	floorHandler := v1.FloorV1Handler{UseCase: useCase}

	r.HandleFunc("/floor-partner/new", floorHandler.CreatePartner).Methods("POST")
	r.HandleFunc("/floor-partners", floorHandler.FindPartners).Methods("GET")
}
