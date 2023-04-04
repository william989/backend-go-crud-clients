package service

import (
	"errors"
	"golang-crud-rest-api/entity"
	"golang-crud-rest-api/repository"
	"log"
	"math/rand"
)

type ClientService interface {
	Validate(cliente *entity.Client) error
	Create(cliente *entity.Client) (*entity.Client, error)
	Update(cliente *entity.Client) (*entity.Client, error)
	FindAll() ([]entity.Client, error)
	FindClienteById(clienteID string) (entity.Client, error)
}

type service struct{}

var (
	repo repository.ClientRepository
)

func NewClientService(repositorio repository.ClientRepository) ClientService {
	repo = repositorio
	return &service{}
}

func (*service) Validate(cliente *entity.Client) error {
	if cliente == nil {
		err := errors.New("Cliente Vacio")
		return err
	}
	if cliente.Nombres == "" {
		err := errors.New("Nombre del Cliente Vacio")
		return err
	}
	return nil
}

// Create implements ClientService
func (*service) Create(cliente *entity.Client) (*entity.Client, error) {

	cliente.ID = rand.Int63()
	return repo.Save(cliente)
}

// Create implements ClientService
func (*service) Update(cliente *entity.Client) (*entity.Client, error) {
	return repo.Actualizar(cliente)
}

// FindById implements ClientService
func (*service) FindClienteById(clienteID string) (entity.Client, error) {
	log.Printf("CONTROLLER - clientId: %v ", clienteID)
	return repo.FindById(clienteID)
}

// FindAll implements ClientService
func (*service) FindAll() ([]entity.Client, error) {
	return repo.FindAll()
}
