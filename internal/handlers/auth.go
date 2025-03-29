package handlers

import (
	"encoding/json"
	"net/http"

	"your_project/config"
	"your_project/internal/models"
	"your_project/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	config.DB.Create(&user)
	w.Write([]byte("User registered successfully"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	var dbUser models.User
	config.DB.Where("email = ?", user.Email).First(&dbUser)

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateJWT(user.Email)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
