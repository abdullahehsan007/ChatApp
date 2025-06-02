package main

import (
	authservice "chatapp/api/auth_service"
	"chatapp/db"
	"chatapp/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConn := db.Database()
	defer dbConn.Close()

	dbLayer := db.NewUserRepository(dbConn)
	serviceLayer := authservice.NewAuthService(dbLayer)
	routerLayer := router.NewRouter(serviceLayer,serviceLayer)

	r := gin.Default()
	routerLayer.RoutersSetup(r)

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
