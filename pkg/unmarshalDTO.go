package pkg

import (
	"customer_partner_match/pkg/pkgError"
	"customer_partner_match/pkg/webResponse"
	"encoding/json"
	"gopkg.in/validator.v2"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type dtoFieldInfo struct {
	name            string
	jsonName        string
	isEmpty         bool
	dtoUnmarshalTag string
}

func UnmarshalDto(w http.ResponseWriter, r *http.Request, dto interface{}) ([]string, pkgError.AppError) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		appError := pkgError.NewInputError("Request DTO error", err)
		webResponse.ERROR(w, http.StatusBadRequest, appError)
		return nil, appError
	}
	appError := unmarshalBodyToDto(w, body, dto)
	if appError != nil {
		return nil, appError
	}
	requestBody := map[string]interface{}{}
	err = json.NewDecoder(strings.NewReader(string(body))).Decode(&requestBody)
	if err != nil {
		appError = pkgError.NewInputError("Request DTO error", err)
		webResponse.ERROR(w, http.StatusBadRequest, appError)
		return nil, appError
	}

	var result []dtoFieldInfo
	dtoReflect := reflect.ValueOf(dto).Elem()
	for i := 0; i < dtoReflect.Type().NumField(); i++ {
		info := dtoFieldInfo{
			name:            dtoReflect.Type().Field(i).Name,
			isEmpty:         IsEmptyValue(dtoReflect.Field(i)),
			dtoUnmarshalTag: dtoReflect.Type().Field(i).Tag.Get("dto-unmarshal"),
			jsonName:        dtoReflect.Type().Field(i).Tag.Get("json"),
		}
		result = append(result, info)
	}

	var updatedFields []string
	for _, fieldInfo := range result {
		if _, ok := requestBody[fieldInfo.jsonName]; ok {
			updatedFields = append(updatedFields, fieldInfo.name)
		}
	}

	return updatedFields, nil
}

func unmarshalBodyToDto(w http.ResponseWriter, body []byte, dto interface{}) pkgError.AppError {
	err := json.Unmarshal(body, dto)
	if err != nil {
		appError := pkgError.NewInputError("Request DTO error", err)
		webResponse.ERROR(w, http.StatusBadRequest, appError)
		return appError
	}

	if err := validator.Validate(dto); err != nil {
		appError := pkgError.NewInputError("Request DTO error", err)
		webResponse.ERROR(w, http.StatusBadRequest, appError)
		return appError
	}
	return nil
}