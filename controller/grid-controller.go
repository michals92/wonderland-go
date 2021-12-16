package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michals92/wonderland-go/errors"
	"github.com/michals92/wonderland-go/service"
)

type gridController struct{}

type GridController interface {
	GetParcels(response http.ResponseWriter, request *http.Request)
}

var gridService service.GridService

func NewPageController(gridSvc service.GridService) GridController {
	gridService = gridSvc
	return &gridController{}
}

func (*gridController) GetParcels(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(request)
	cardId := vars["cardId"]

	fmt.Println(cardId)

	//page, err := pageService.GetPage(cardId)

	//	if err != nil {
	//	sendJson(response, http.StatusBadRequest, "Unable to get page")
	//}

	//response.WriteHeader(http.StatusOK)
	//	json.NewEncoder(response).Encode(page)
	sendJson(response, http.StatusBadRequest, "Get parcels not implemented")
}

func sendJson(response http.ResponseWriter, statusCode int, message string) {
	response.WriteHeader(statusCode)
	json.NewEncoder(response).Encode(errors.ServiceError{Message: message})
}
