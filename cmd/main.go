// cmd/main.go
package main

import (
	"myapp/internal/app/delivery/http"
	"myapp/internal/app/repository"
	"myapp/internal/app/usecase"
	"net/http"
)

func main() {
	// Configura la inyección de dependencias
	userRepository := repository.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := http.NewUserHandler(userUsecase)

	// Configura las rutas HTTP
	http.HandleFunc("/users/{id}", userHandler.GetUserByID)

	// Inicia el servidor HTTP
	http.ListenAndServe(":8080", nil)
}
