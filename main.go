package main

import (
	"fmt"
	"log"
	"net/http"

	"exemplo.com/meuprojeto/entities"
)

func main() {
	http.HandleFunc("/users", entities.UsersHandler)
	http.HandleFunc("/users/", entities.UserHandler)

	http.HandleFunc("/items", entities.ItemsHandler)
	http.HandleFunc("/items/", entities.ItemHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}