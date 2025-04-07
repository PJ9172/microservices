package handlers

import (
	"Microservices/config"
	"Microservices/models"
	"Microservices/utils"
	"fmt"
	"net/http"

)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		r.ParseForm()
		email := r.FormValue("email")
		password := r.FormValue("password")
	
		user := models.User{Email: email, Password: password}
		if err := config.DB.Create(&user).Error; err != nil {
			http.Error(w, "Email already in use", http.StatusBadRequest)
			return
		}
		fmt.Println("User registered!!")
	
		utils.SendWelcomeEmail(user.Email)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Post"{
		email := r.FormValue("email")
		password := r.FormValue("password")
	
		var user models.User
		if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}
		if err := config.DB.Where("password = ?", password).First(&user).Error; err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}
	
		w.Write([]byte("Logged in successfully!"))
	}
}
