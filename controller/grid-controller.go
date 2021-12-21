package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/michals92/wonderland-go/entity"
	"github.com/michals92/wonderland-go/errors"
	"github.com/michals92/wonderland-go/service"
)

type gridController struct{}

type GridController interface {
	GetParcels(response http.ResponseWriter, request *http.Request)
	AddParcel(response http.ResponseWriter, request *http.Request)
	PinArt(response http.ResponseWriter, request *http.Request)
	UnpinArt(response http.ResponseWriter, request *http.Request)
}

var gridService service.GridService

func NewGridController(gridSvc service.GridService) GridController {
	gridService = gridSvc
	return &gridController{}
}

func (*gridController) GetParcels(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json; charset=utf-8")

	decoder := json.NewDecoder(request.Body)
	var userInfo entity.UserInfo
	error := decoder.Decode(&userInfo)

	if error != nil {
		sendJson(response, http.StatusBadRequest, "Unable to parse bounding box")
		return
	}

	parcels, error := gridService.GetGrid(&userInfo)
	if error != nil {
		sendJson(response, http.StatusBadRequest, error.Error())
		return
	}

	fmt.Println(parcels)

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(parcels)
}

func (*gridController) AddParcel(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json; charset=utf-8")

	decoder := json.NewDecoder(request.Body)
	var parcel entity.Parcel
	error := decoder.Decode(&parcel)

	if error != nil {
		sendJson(response, http.StatusBadRequest, "Unable to parse parcel")
		return
	}

	error = gridService.AddParcel(&parcel)
	if error != nil {
		sendJson(response, http.StatusBadRequest, error.Error())
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (*gridController) PinArt(response http.ResponseWriter, request *http.Request) {

}

func (*gridController) UnpinArt(response http.ResponseWriter, request *http.Request) {

}

func sendJson(response http.ResponseWriter, statusCode int, message string) {
	response.WriteHeader(statusCode)
	json.NewEncoder(response).Encode(errors.ServiceError{Message: message})
}
