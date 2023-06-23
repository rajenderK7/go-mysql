package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rajenderK7/go-mysql/pkg/database"
	"github.com/rajenderK7/go-mysql/pkg/routes"
	"github.com/rajenderK7/go-mysql/pkg/utils"
)

var (
	DBName     string
	DBUser     string
	DBPassword string
	DBPORT     string
)

func loadENV() {
	err := godotenv.Load("../../.env")
	utils.CheckGeneralErr(err)
	DBName = os.Getenv("DB_NAME")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBPORT = os.Getenv("DB_PORT")
}

func main() {
	app := *fiber.New()
	loadENV()
	config := database.DBConfig{
		DBName:     DBName,
		DBUser:     DBUser,
		DBPassword: DBPassword,
		DBPORT:     DBPORT,
	}
	database.ConnectDB(&config)
	database.RunMigrations()
	routes.SetupRoutes(&app)
	sqlBD, _ := database.GetDB().DB()
	defer sqlBD.Close()
	log.Fatal(app.Listen(":4000"))
}
