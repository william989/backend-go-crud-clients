package repository

import (
	"context"
	"log"

	"golang-crud-rest-api/entity"

	"cloud.google.com/go/firestore"
)

type ClientRepository interface {
	Save(cliente *entity.Client) (*entity.Client, error)
	FindAll() ([]entity.Client, error)
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
		log.Print(clientes)

	}
	return clientes, nil
}
