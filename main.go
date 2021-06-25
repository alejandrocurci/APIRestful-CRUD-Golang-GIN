package main

import (
	"api-crud-gin/server"
)

func main() {

	server := server.Build()
	server.Run(":8080")

}
