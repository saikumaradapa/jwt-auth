package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/saikumaradapa/jwt-auth/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	var user models.User

	errDecode := json.NewDecoder(r.Body).Decode(&user)
	if errDecode != nil {
		log.Fatalf("error while decoding %v", errDecode)
	}

	models.Users[user.Username] = user

	w.WriteHeader(http.StatusCreated)
	successMsg := map[string]string{"message": "user registered"}
	errEncode := json.NewEncoder(w).Encode(successMsg)
	if errEncode != nil {
		log.Fatalf("error while decoding %v", errEncode)
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	var credentials models.User
	errDecode := json.NewDecoder(r.Body).Decode(&credentials)
	if errDecode != nil {
		log.Fatalf("error while decoding %v", errDecode)
	}

	user, exists := models.Users[credentials.Username]
	if !exists || user.Password != credentials.Password {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(credentials.Username)
	if err != nil {
		http.Error(w, "could not generate token", http.StatusInternalServerError)
		return
	}

	tokenMsg := map[string]string{"token": token}
	errEncode := json.NewEncoder(w).Encode(tokenMsg)
	if errEncode != nil {
		log.Fatalf("error while decoding %v", errEncode)
	}
}

func Protected(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	welcomeMsg := map[string]string{"message": fmt.Sprintf("Welcome %v", username)}
	errEncode := json.NewEncoder(w).Encode(welcomeMsg)
	if errEncode != nil {
		log.Fatalf("error while decoding %v", errEncode)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users := make([]models.User, 0, len(models.Users))
	for _, user := range models.Users {
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
