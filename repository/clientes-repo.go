package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"golang-crud-rest-api/entity"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClientRepository interface {
	Save(cliente *entity.Client) (*entity.Client, error)
	Actualizar(cliente *entity.Client) (*entity.Client, error)
	FindAll() ([]entity.Client, error)
	FindById(clienteID string) (entity.Client, error)
}

type repo struct {
}

// NewClientRepository
func NewClientRepository() ClientRepository {
	return &repo{}
}

const (
	projectId      string = "desafio-rest-crud-clientes"
	collectionName string = "clientes"
)

func (*repo) Save(cliente *entity.Client) (*entity.Client, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Falla al crear Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":      cliente.ID,
		"Nombres": cliente.Nombres,
		"Email":   cliente.Email,
	})

	if err != nil {
		log.Fatalf("Falla al agregar nuevo cliente: %v", err)
		return nil, err
	}
	return cliente, nil
}

func (*repo) Actualizar(cliente *entity.Client) (*entity.Client, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Falla al crear Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	docRef := client.Collection(collectionName).Doc("ID")

	_, err = docRef.Update(ctx, []firestore.Update{
		{Path: "Nombres", Value: cliente.Nombres},
		{Path: "Email", Value: cliente.Email},
	})

	if err != nil {
		log.Fatalf("Falla al agregar nuevo cliente: %v", err)
		return nil, err
	}
	return cliente, nil
}

func (*repo) FindById(clienteID string) (entity.Client, error) {

	// Inicializar el cliente de Firestore
	ctx := context.Background()
	var cliente entity.Client
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		return cliente, err
	}
	defer client.Close()

	// Obtener una referencia al documento con el ID especificado
	docRef := client.Collection("clients").Doc(clienteID)

	// Obtener los datos del documento
	doc, err := docRef.Get(ctx)
	if err != nil {
		log.Print(cliente.Nombres)
		log.Print(err)
		if status.Code(err) == codes.NotFound {
			return cliente, fmt.Errorf("Cliente con ID %s no encontrado", clienteID)
		}
		return cliente, err
	}
	if err := doc.DataTo(&cliente); err != nil {
		return cliente, err
	}
	cliente.ID, _ = strconv.ParseInt(docRef.ID, 10, 64)
	return cliente, nil
}

func (*repo) FindAll() ([]entity.Client, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Falla al crear Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	var clientes []entity.Client
	iterator := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iterator.Next()
		if doc == nil {
			break
		}
		if err != nil {
			log.Fatalf("Falla al iterar Lista de clientes: %v", err)
			return nil, err

		}
		cliente := entity.Client{
			ID:      doc.Data()["ID"].(int64),
			Nombres: doc.Data()["Nombres"].(string),
			Email:   doc.Data()["Email"].(string),
		}
		clientes = append(clientes, cliente)

	}
	return clientes, nil
}
