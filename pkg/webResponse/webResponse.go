package webResponse

import (
	"customer_partner_match/pkg/pkgError"
	"encoding/json"
	"fmt"
	"net/http"
)

type successResponse struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Err      string `json:"err"`
	Message  string `json:"message"`
	ErrorKey string `json:"errorKey"`
	Status   int    `json:"status"`
}

func ERROR(w http.ResponseWriter, statusCode int, err pkgError.AppError) {
	errResponse := errorResponse{
		ErrorKey: err.GetErrorKey(),
		Message:  err.GetMessage(),
		Err:      err.GetErr(),
		Status:   statusCode,
	}

	w.WriteHeader(statusCode)

	build(w, statusCode, errResponse)
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	if data != nil {
		response := successResponse{
			Data: data,
		}
		build(w, statusCode, response)
		return
	}

	build(w, statusCode, nil)
}

func build(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
	}
}
