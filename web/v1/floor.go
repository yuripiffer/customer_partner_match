package v1

import (
	"customer_partner_match/model"
	"customer_partner_match/pkg"
	"customer_partner_match/pkg/pkgError"
	"customer_partner_match/pkg/webResponse"
	"customer_partner_match/ports/input"
	"errors"
	"math"
	"net/http"
)

type FloorV1Handler struct {
	UseCase input.FloorUseCase
}

func (h *FloorV1Handler) CreatePartner(w http.ResponseWriter, r *http.Request) {
	requestDTO := model.NewFloorPartnerDTO{}
	_, appError := pkg.UnmarshalDto(w, r, &requestDTO)
	if appError != nil {
		return
	}

	missingFields := checkNewPartnerDTO(requestDTO)
	if missingFields != "" {
		webResponse.ERROR(w, http.StatusBadRequest,
			pkgError.NewInputError("missing/invalid field(s)", errors.New(missingFields)))
		return
	}

	partner, appError := h.UseCase.CreatePartner(r.Context(), requestDTO)
	if appError != nil {
		//TODO treat errors
		return
	}
	webResponse.JSON(w, http.StatusOK, partner)
	return
}

func (h *FloorV1Handler) FindPartners(w http.ResponseWriter, r *http.Request) {
	requestDTO := model.FloorRequestDTO{}
	_, appError := pkg.UnmarshalDto(w, r, &requestDTO)
	if appError != nil {
		return
	}

	missingFields := checkFindPartnersDTO(requestDTO)
	if missingFields != "" {
		webResponse.ERROR(w, http.StatusBadRequest,
			pkgError.NewInputError("missing/invalid field(s)", errors.New(missingFields)))
		return
	}

	floorPartners, appError := h.UseCase.FindPartners(r.Context(), requestDTO)
	if appError != nil {
		//TODO treat errors
		return
	}
	webResponse.JSON(w, http.StatusOK, floorPartners)
	return
}

func checkNewPartnerDTO(requestDTO model.NewFloorPartnerDTO) string {
	missingFields := ""
	if math.Abs(requestDTO.Latitude) > 180. {
		missingFields += "latitude, "
	}
	if math.Abs(requestDTO.Longitude) > 180. {
		missingFields += "longitude, "
	}
	if requestDTO.Partner == "" {
		missingFields += "partner, "
	}
	if requestDTO.OperatingRadius <= 0 {
		missingFields += "operating_radius, "
	}

	if (requestDTO.Wood && requestDTO.Tiles && requestDTO.Carpet) == false {
		missingFields += "no material informed (wood, tiles and/or carpet), "
	}
	if missingFields != "" {
		return missingFields[:len(missingFields)-2]
	}
	return ""
}

func checkFindPartnersDTO(requestDTO model.FloorRequestDTO) string {
	missingFields := ""
	if math.Abs(requestDTO.Latitude) > 180. {
		missingFields += "latitude, "
	}
	if math.Abs(requestDTO.Longitude) > 180. {
		missingFields += "longitude, "
	}
	if requestDTO.FloorArea <= 0 {
		missingFields += "floor_area, "
	}
	if requestDTO.Phone == "" {
		missingFields += "phone, "
	}
	if (requestDTO.Wood && requestDTO.Tiles && requestDTO.Carpet) == false {
		missingFields += "no material selected (wood, tiles and/or carpet), "
	}
	if missingFields != "" {
		return missingFields[:len(missingFields)-2]
	}
	return ""
}
