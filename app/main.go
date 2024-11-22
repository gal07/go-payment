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

	// Load ENV
	basepath, err := os.Getwd()
	if err != nil {
		return
	}

	pth := filepath.Join(basepath + "/.envs")
	fmt.Println(pth)
	errs := godotenv.Load(pth)
	if errs != nil {
		log.Fatal("Error loading .env file")
	}

	// server stuff
	serverPort := os.Getenv("SERVER_PORT")
	fmt.Println(os.Getenv("SERVER_PORT"))
	fmt.Println(os.Getenv("DATABASE_NAME"))

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
