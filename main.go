package main

import (
	"github.com/ade-iqbal/fga-assignment-2/database"
	"github.com/ade-iqbal/fga-assignment-2/routes"
)

func main() {
	PORT := ":8080"

	defer database.CloseConnection()
	routes.StartServer().Run(PORT)
}