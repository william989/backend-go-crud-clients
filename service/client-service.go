package service

import (
	"errors"
	"golang-crud-rest-api/entity"
	"golang-crud-rest-api/repository"
	"math/rand"
)

type ClientService interface {
	Validate(cliente *entity.Client) error
	Create(cliente *entity.Client) (*entity.Client, error)
	FindAll() ([]entity.Client, error)
}

type service struct{}

var (
	repo repository.ClientRepository = repository.NewClientRepository()
)

func NewClientService() ClientService {
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

// FindAll implements ClientService
func (*service) FindAll() ([]entity.Client, error) {
	return repo.FindAll()
}
