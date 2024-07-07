package repository

import (
	"summer_practice/internal/domain"
	database "summer_practice/internal/jsonDatabase"
)

type Repository struct {
	db *database.DataBase
}

func New(db *database.DataBase) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) CreateCarRepository(car domain.Car) (domain.Car, error) {
	row, err := r.db.AddRow(car)
	if err != nil {
		return domain.Car{}, err
	}

	return row, nil
}
func (r Repository) GetAllCarsRepository() ([]domain.Car, error) {
	rows, err := r.db.GetAllRows()
	if err != nil {
		return nil, err
	}
	return rows, nil
}
func (r Repository) GetCarByIdRepository(id int) (domain.Car, error) {
	row, err := r.db.GetRow(id)
	if err != nil {
		return domain.Car{}, err
	}
	return row, nil
}
func (r Repository) PutCarByIdRepository(car domain.Car) (domain.Car, error) {
	car, err := r.db.PutRow(car)
	if err != nil {
		return domain.Car{}, err
	}
	return car, nil
}
func (r Repository) PatchCarByIdRepository(fieldsToUpdate map[string]interface{}) (domain.Car, error) {
	car, err := r.db.UpdateRow(fieldsToUpdate)
	if err != nil {
		return domain.Car{}, err
	}
	return car, nil
}
func (r Repository) DeleteCarRepository(id int) error {
	if err := r.db.DeleteRow(id); err != nil {
		return err
	}
	return nil
}
