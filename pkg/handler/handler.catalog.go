package handler

import (
	"encoding/json"
	"github.com/alisyahbana/bcg-test/pkg/service/catalog"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"message"`
}

type MessageResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func PurchaseHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	errorResponse := ErrorResponse{}
	jsonBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Header().Set("Content-Type", "application/json")
		errorResponse.Error = err.Error()
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}

	var payload catalog.PurchasePayload
	err = json.Unmarshal(jsonBody, &payload)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Header().Set("Content-Type", "application/json")
		errorResponse.Error = err.Error()
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}

	purchaseResponse, err := catalog.New().Purchase(payload)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Header().Set("Content-Type", "application/json")
		errorResponse.Error = err.Error()
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}

	message := MessageResponse{
		Message: "Success",
		Data:    purchaseResponse,
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(message)
}
