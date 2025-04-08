package handlers

import (
	"Microservices/config"
	"Microservices/models"
	"Microservices/utils"
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		email := r.FormValue("email")
		password := r.FormValue("password")

		user := models.User{Email: email, Password: password}
		if err := config.DB.Create(&user).Error; err != nil {
			http.Error(w, "Email already in use!!!\nPlease SignIn insted of SignUp!!!", http.StatusBadRequest)
			return
		}
		fmt.Println("User registered!!")

		utils.SendWelcomeEmail(user.Email)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		email := r.FormValue("email")
		password := r.FormValue("password")

		var user models.User
		if err := config.DB.Where("email = ? and password = ?", email, password).First(&user).Error; err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		session, _ := config.Store.Get(r, "session")
		session.Values["email"] = user.Email
		session.Save(r, w)
		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func ShowProducts(w http.ResponseWriter, r *http.Request) {
	products := []string{"Apples", "Banana", "Mango"}

	tml := template.Must(template.ParseFiles("templates/products.html"))
	tml.Execute(w, products)
}

func PurchaseProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		session, _ := config.Store.Get(r, "session")
		email := session.Values["email"].(string)

		r.ParseForm()
		products := r.Form["products"]

		rawData := strings.Join(products, ",")

		purchase := models.Purchase{
			Email:   email,
			RawData: rawData,
			Status:  "Pending",
		}

		config.DB.Create(&purchase)
		tmpl := template.Must(template.ParseFiles("templates/payment.html"))
		tmpl.Execute(w, nil)
	}
}

func MakePayment(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, "session")
	email := session.Values["email"].(string)

	var purchase models.Purchase
	config.DB.Where("email = ? and status = ?", email, "Pending").First(&purchase)

	purchase.Status = "Paid"
	config.DB.Save(&purchase)
	utils.SendPaymentEmail(email)

	tmpl := template.Must(template.ParseFiles("templates/success.html"))
	tmpl.Execute(w, nil)
}
