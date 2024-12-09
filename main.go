package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mauri247/db-api/db"
	"github.com/mauri247/db-api/handlers"
)

func main() {
	// Leer la URI de MongoDB desde las variables de entorno
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("La variable de entorno MONGO_URI no está configurada")
	}

	// Conectar a MongoDB
	if err := db.ConnectMongo(mongoURI); err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}

	// Registrar los endpoints
	http.HandleFunc("/characters", handlers.AddCharacterHandler)
	http.HandleFunc("/characters/search", handlers.GetCharactersByNameHandler)

	// Iniciar el servidor
	log.Println("Servidor corriendo en el puerto 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
