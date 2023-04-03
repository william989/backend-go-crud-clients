package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"golang-crud-rest-api/entity"
	"golang-crud-rest-api/service"
)

var (
	clientService service.ClientService = service.NewClientService()
)

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	clientes, err := clientService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.New("Error al obtener clientes"))

	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(clientes)

}

func addPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var client entity.Client
	err := json.NewDecoder(req.Body).Decode(&client)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.New("Error al agregar clientes"))
		return
	}
	err1 := clientService.Validate(&client)
	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.New("Error al X agregar clientes"))
		return
	}

	clientService.Create(&client)

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(client)
}
