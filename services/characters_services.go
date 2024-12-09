package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/mauri247/db-api/db"
	"github.com/mauri247/db-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	dbName               = "dragonball-db"
	charactersCollection = "characters"
)

func AddCharacter(name string) error {
	character, err := GetCharactersByName(name)
	if err != nil {
		fmt.Println("Error al consultar en MongoDB")
		return err
	}
	if len(character) > 0 {
		var miError error = errors.New("El personaje ya existe en la base de datos, intente con otro nombre")
		return miError
	}

	characterData, err := GetCharacterData(name)
	if err != nil {
		fmt.Println("Error al obtener los datos del personaje")
		return err
	}

	newCharacter := models.Character{
		Name: characterData.Name,
		Race: characterData.Race,
		Ki:   characterData.Ki,
	}

	collection := db.Client.Database(dbName).Collection(charactersCollection)
	_, err = collection.InsertOne(context.TODO(), newCharacter)
	return err
}

func GetCharactersByName(name string) ([]models.Character, error) {
	collection := db.Client.Database(dbName).Collection(charactersCollection)
	filter := bson.M{"name": bson.M{"$regex": `(?i)^` + name + `$`}}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var characters []models.Character
	for cursor.Next(context.TODO()) {
		var character models.Character
		if err := cursor.Decode(&character); err != nil {
			return nil, err
		}
		characters = append(characters, character)
	}

	return characters, nil
}

func GetCharacterData(name string) (*models.CharacterApi, error) {

	baseURL := "https://dragonball-api.com/api/characters"
	encodedName := url.QueryEscape(name)
	url := fmt.Sprintf("%s?name=%s", baseURL, encodedName)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error al hacer la solicitud:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error al leer la respuesta:", err)
		return nil, err
	}

	var characters []models.CharacterApi
	if err := json.Unmarshal(body, &characters); err != nil {
		log.Println("Error al decodificar el JSON:", err)
		return nil, err
	}

	if len(characters) == 0 {
		return nil, fmt.Errorf("no se encontraron personajes con el nombre: %s en la api externa", name)
	}

	return &characters[0], nil
}
