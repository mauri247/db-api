package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mauri247/db-api/models"
	"github.com/mauri247/db-api/services"
)

func AddCharacterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var payload models.PayloadData
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Error al procesar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	var name string
	if payload.Name != "" {
		name = payload.Name
	} else if payload.Character != "" {
		name = payload.Character
	} else {
		errorResponse := map[string]interface{}{
			"Ok":          false,
			"Description": "Faltan los campos 'name' o 'character'",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if err := services.AddCharacter(name); err != nil {
		errorResponse := map[string]interface{}{
			"Ok":          false,
			"Description": err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	successResponse := map[string]interface{}{
		"OK":          true,
		"Description": "Personaje guardado exitosamente",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(successResponse)

}

func GetCharactersByNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse := map[string]interface{}{
			"Ok":          false,
			"Description": "Método no permitido",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		errorResponse := map[string]interface{}{
			"Ok":          false,
			"Description": "El parámetro 'name' es obligatorio",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	characters, err := services.GetCharactersByName(name)
	if err != nil {
		errorResponse := map[string]interface{}{
			"Ok":          false,
			"Description": "El parámetro 'name' es obligatorio",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	successResponse := map[string]interface{}{
		"OK":   true,
		"Data": characters,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(successResponse)
}
