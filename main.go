package main

import (
	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
	"dumbmerch/routes"
	"fmt"
	"os"

	"github.com/joho/godotenv" // import this package
	"github.com/labstack/echo"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	e := echo.New()

	mysql.DatabaseInit()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	e.Static("/uploads", "./uploads")

	PORT := os.Getenv("PORT")

	fmt.Println("server running on port" + PORT)
	e.Logger.Fatal(e.Start(":" + PORT))
}
