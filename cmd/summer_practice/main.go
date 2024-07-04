package main

import (
	"log"
	database "summer_practice/pkg"
	"time"
)

//func init() {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//}

func main() {
	db := database.New("data.json")
	//car1 := domain.Car{
	//	ID:          1337,
	//	Brand:       "Tesla",
	//	Model:       "X",
	//	Mileage:     100,
	//	OwnersCount: 1,
	//}
	//car2 := domain.Car{
	//	ID:          1338,
	//	Brand:       "Toyota",
	//	Model:       "Mark2",
	//	Mileage:     100,
	//	OwnersCount: 1,
	//}

	rows, err := db.GetAllRows()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(rows)
	time.Sleep(20 * time.Second)

}
