package main

import (
	"fmt"
	"os"

	"github.com/alwialdi9/be_auth-jajanskuy/connection"
	"github.com/alwialdi9/be_auth-jajanskuy/routes"
	"github.com/alwialdi9/be_auth-jajanskuy/utils"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found, using environment variables")
	}

	log.SetReportCaller(true)
	log.SetFormatter(&utils.Formatter{})
	log.SetReportCaller(true)

	connection.NewConnection()
	connection.CheckConnection()

	app := routes.Router()
	if err := app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Info("Server is running on port 3000")
	log.Info("Visit http://localhost:3000/api/health to check the server status")
}
