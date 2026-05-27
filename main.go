package main

import (
	"fmt"
	libraryhttp "homework/http"
	"homework/library"
)

func main() {
	l := library.NewLibrary()

	handlers := libraryhttp.NewHandlers(l)
	server := libraryhttp.NewServer(handlers)

	err := server.StartServer()
	if err != nil {
		fmt.Println("error while start server", err)
	}
}
