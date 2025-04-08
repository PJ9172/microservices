package main

import (
	"Microservices/config"
	"Microservices/models"
	"Microservices/routers"
	"log"
	"net/http"
	"os"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Purchase{})

	r := routers.SetRouters()
	port := os.Getenv("PORT")

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
