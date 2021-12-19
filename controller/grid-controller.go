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
}

var gridService service.GridService

func NewGridController(gridSvc service.GridService) GridController {
	gridService = gridSvc
	return &gridController{}
}

func (*gridController) GetParcels(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json; charset=utf-8")

	decoder := json.NewDecoder(request.Body)
	var boundingBox entity.BoundingBox
	error := decoder.Decode(&boundingBox)

	if error != nil {
		sendJson(response, http.StatusBadRequest, "Unable to parse bounding box")
	}

	parcels, error := gridService.GetGrid(&boundingBox)
	if error != nil {
		sendJson(response, http.StatusBadRequest, "Error getting grid")
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
		sendJson(response, http.StatusBadRequest, "Unable to add parcel")
		return
	}

	response.WriteHeader(http.StatusOK)
}

func sendJson(response http.ResponseWriter, statusCode int, message string) {
	response.WriteHeader(statusCode)
	json.NewEncoder(response).Encode(errors.ServiceError{Message: message})
}
