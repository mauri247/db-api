package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongo(uri string) error {
	var err error
	Client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("error al crear cliente de MongoDB: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = Client.Connect(ctx)
	if err != nil {
		return fmt.Errorf("error al conectar a MongoDB: %w", err)
	}

	err = Client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("error al hacer ping a MongoDB: %w", err)
	}

	fmt.Println("Conexión a MongoDB exitosa")
	return nil
}
