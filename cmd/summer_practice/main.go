package main

import (
	"summer_practice/internal/handler"
	database "summer_practice/internal/jsonDatabase"
	"summer_practice/internal/repository"
	"summer_practice/internal/service"
)

func main() {
	db := database.New("data.json")
	repo := repository.New(db)
	svc := service.New(repo)
	h := handler.New(svc)

	h.Run(":8080")
}
