package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"golang-crud-rest-api/entity"
	"golang-crud-rest-api/service"

	"github.com/gorilla/mux"
)

type controller struct{}

var (
	clientService service.ClientService
)

type ClientController interface {
	GetClientes(resp http.ResponseWriter, req *http.Request)
	AddCliente(resp http.ResponseWriter, req *http.Request)
	GetClienteById(resp http.ResponseWriter, req *http.Request)
}

func NewClientController(service service.ClientService) ClientController {
	clientService = service
	return &controller{}
}

func (*controller) GetClientes(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	clientes, err := clientService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.New("Error al obtener clientes"))

	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(clientes)

}

func (*controller) AddCliente(resp http.ResponseWriter, req *http.Request) {
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
		json.NewEncoder(resp).Encode(errors.New("Error al validar clientes"))
		return
	}

	clientService.Create(&client)

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(client)
}

func (*controller) UpdateCliente(resp http.ResponseWriter, req *http.Request) {

	resp.WriteHeader(http.StatusOK)
}

func (*controller) GetClienteById(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	clientId := mux.Vars(req)["ID"]
	log.Print("CONTROLLER - clientId: " + clientId)
	/*params := mux.Vars(req)
	clientId, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.New("Error al obtener el ID del cliente"))
		return
	}*/

	cliente, err1 := clientService.FindClienteById(clientId)
	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.New("Error al obtener el cliente"))
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(cliente)
}
