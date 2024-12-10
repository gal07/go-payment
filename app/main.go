package main

import (
	"fmt"
	"go-payment/middleware"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	directory, err := filepath.Abs(filepath.Dir(os.Args[0])) //get the current working directory
	if err != nil {
		log.Fatal(err) //print the error if obtained
	}
	log.Println("Current working directory:", directory)

	errs := godotenv.Load(directory + "/.env")
	if errs != nil {
		log.Fatal("Error loading .env file")
	}

	// server stuff
	serverPort := os.Getenv("GP_SERVER_PORT")
	fmt.Println(os.Getenv("GP_SERVER_PORT"))
	fmt.Println(os.Getenv("GP_DATABASE_NAME"))

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
