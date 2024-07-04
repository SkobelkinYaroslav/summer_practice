package repository

import (
	"summer_practice/internal/domain"
	database "summer_practice/pkg"
)

type Repository struct {
	db *database.DataBase
}

func (r *Repository) CreateCarRepository(car domain.Car) (domain.Car, error){
	r.db.AddRow(car)
}
func (r *Repository) GetAllCarsRepository() ([]domain.Car, error){
	r.db.GetRow()
}
func (r *Repository) GetCarByIdRepository(id int) (domain.Car, error){
	r.db.
}
func (r *Repository) UpdateCarByIdRepository(car domain.Car) (domain.Car, error){
	r.db.
}
func (r *Repository) PatchCarByIdRepository(car domain.Car) (domain.Car, error){
	r.db.
}
func (r *Repository) DeleteCarRepository(id int) error{
	r.db.
}
