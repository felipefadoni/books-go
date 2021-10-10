package main

import (
	"go-gin-gonic/src/database"
	"go-gin-gonic/src/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
