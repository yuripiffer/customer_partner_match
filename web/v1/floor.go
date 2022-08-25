package v1

import (
	"customer_partner_match/pkg"
	"customer_partner_match/pkg/pkgError"
	"customer_partner_match/pkg/webResponse"
	"customer_partner_match/ports/input"
	"errors"
	"math"
	"net/http"
)

type FloorV1Handler struct {
	useCase input.FloorUseCase
}

func (h *FloorV1Handler) NewPartner(w http.ResponseWriter, r *http.Request) {}

func (h *FloorV1Handler) FindPartners(w http.ResponseWriter, r *http.Request) {
	requestDTO := input.FloorRequestDTO{}
	_, appError := pkg.UnmarshalDto(w, r, &requestDTO)
	if appError != nil {
		return
	}
	if math.Abs(requestDTO.Latitude) > 180 || math.Abs(requestDTO.Longitude) > 180 {
		webResponse.ERROR(w, http.StatusBadRequest,
			pkgError.NewInputError("", errors.New("invalid latitude/longitude")))
		return
	}

	missingFields := ""
	if (requestDTO.Wood && requestDTO.Tiles && requestDTO.Carpet) == false {
		missingFields += "no material selected, "
	}
	if requestDTO.FloorArea == 0 {
		missingFields += "missing floor area (square meter), "
	}
	if requestDTO.Phone == "" {
		missingFields += "missing phone, "
	}

	if missingFields != "" {
		webResponse.ERROR(w, http.StatusBadRequest,
			pkgError.NewInputError("", errors.New(missingFields[:len(missingFields)-2])))
		return
	}

	appError = h.useCase.FindPartner(r.Context(), requestDTO)
	if appError != nil {
		//TODO treat errors
	}
	webResponse.JSON(w, http.StatusOK, nil)

}
