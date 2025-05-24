package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/saikumaradapa/jwt-auth/handlers"
)

func main() {
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/users", handlers.GetUsers)
	http.HandleFunc("/protected", middleware.JWTMiddleware(http.HandlerFunc(handlers.Protected)))

	fmt.Println("server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
