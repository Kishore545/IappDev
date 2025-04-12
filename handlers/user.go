package handlers

import (
	"IAppDev/models"
	"encoding/json"
	"net/http"
)

var users = []models.User{}
var idCounter = 1

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	user.ID = idCounter
	idCounter++
	users = append(users, user)

	json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	json.NewDecoder(r.Body).Decode(&creds)

	for _, user := range users {
		if user.Username == creds.Username && user.Password == creds.Password {
			w.Write([]byte("Login successful"))
			return
		}
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}
