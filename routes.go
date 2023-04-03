package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"golang-crud-rest-api/entity"
	"golang-crud-rest-api/repository"
)

var (
	repo repository.ClientRepository = repository.NewClientRepository()
)

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	clientes, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{error: "Error al obtener clientes"}`))
		return
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
		resp.Write([]byte(`{error: "Error unmarshalling the request"}`))
		return
	}
	client.ID = rand.Int63()
	repo.Save(&client)

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(client)
}
