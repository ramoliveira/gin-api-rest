package main

import (
	"github.com/ramoliveira/api-go-gin/database"
	"github.com/ramoliveira/api-go-gin/routes"
)

func main() {
	database.ConnectWithDatabase()
	routes.HandleRequests()
}
