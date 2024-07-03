package main

import (
	"fmt"
	"summer_practice/internal/domain"
	database "summer_practice/pkg"
)

//func init() {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//}

func main() {
	db := database.New("data.json")
	//fmt.Println(db.GetRow(1))
	fmt.Println(db)
	car := domain.Car{
		ID:          1337,
		Brand:       "Tesla",
		Model:       "X",
		Mileage:     100,
		OwnersCount: 1,
	}
	db.AddRow(car)
	fmt.Println(db)
}
