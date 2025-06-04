package main

import (
	authservice "chatapp/api/auth_service"
	mongodb "chatapp/db/MongoDB"
	postgresql "chatapp/db/Postgresql"

	"chatapp/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	
	mongodb.Connect()
	dbConn := postgresql.Database()
	defer dbConn.Close()

	dbLayer := postgresql.NewUserRepository(dbConn)
	serviceLayer := authservice.NewAuthService(dbLayer)
	routerLayer := router.NewRouter(serviceLayer, serviceLayer)

	r := gin.Default()
	routerLayer.RoutersSetup(r)

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
