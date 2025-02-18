package main

import (
	"database/sql"
	"log"
	"net/http"

	"dev_toolbox_auth_service/internal/delivery/handler"
	"dev_toolbox_auth_service/internal/repository/postgres"
	"dev_toolbox_auth_service/internal/usecase"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Подключаемся к БД
	db, err := sql.Open("postgres", "postgres://auth_user:auth_pass@localhost:5432/auth_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Инициализируем слои
	userRepo := postgres.NewUserRepository(db)
	authUseCase := usecase.NewAuthUseCase(userRepo)
	authHandler := handler.NewAuthHandler(authUseCase)

	// Настраиваем роутер
	r := mux.NewRouter()
	r.HandleFunc("/auth/register", authHandler.Register).Methods("POST")

	log.Printf("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
