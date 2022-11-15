package main

import (
	"farms/internal/server"
)

func main() {
	srv := server.NewServer("8080")
	srv.Run()
}
