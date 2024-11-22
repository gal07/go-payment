package main

import (
	"fmt"
	"go-payment/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load ENV
	errs := godotenv.Load(".env")
	if errs != nil {
		log.Fatal("Error loading .env file")
	}

	// server stuff
	serverPort := os.Getenv("GP_SERVER_PORT")
	fmt.Println(os.Getenv("GP_SERVER_PORT"))
	fmt.Println(os.Getenv("GP_XENDIT_API"))
	fmt.Println(os.Getenv("GP_BASIC_AUTH_USERNAME"))

	// release Mode
	// release := os.Getenv("RELEASE_MODE")
	// if release == "release" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	route := gin.Default()
	route.Use(middleware.MiddlewareAuth)

	// Route
	service(route)

	// Run
	route.Run(serverPort)

}
