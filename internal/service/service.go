package service

import (
	"summer_practice/internal/domain"
)

type CarRepository interface {
	CreateCarRepository(car domain.Car) (domain.Car, error)
	GetAllCarsRepository() ([]domain.Car, error)
	GetCarByIdRepository(id int) (domain.Car, error)
	PutCarByIdRepository(car domain.Car) (domain.Car, error)
	PatchCarByIdRepository(fieldsToUpdate map[string]interface{}) (domain.Car, error)
	DeleteCarRepository(id int) error
}

type CarService struct {
	Repository CarRepository
}

func New(repo CarRepository) *CarService {
	return &CarService{
		Repository: repo,
	}
}

func (s *CarService) CreateCarService(car domain.Car) (domain.Car, error) {
	if car.Model == "" || car.Brand == "" || car.Mileage < 0 || car.OwnersCount < 0 {
		return domain.Car{}, domain.ErrBadParamInput
	}
	carOutput, err := s.Repository.CreateCarRepository(car)
	if err != nil {
		return domain.Car{}, err
	}
	return carOutput, nil
}
func (s *CarService) GetAllCarsService() ([]domain.Car, error) {
	carArray, err := s.Repository.GetAllCarsRepository()

	if len(carArray) == 0 {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return carArray, nil
}
func (s *CarService) GetCarByIdService(id int) (domain.Car, error) {
	if id < 0 {
		return domain.Car{}, domain.ErrBadParamInput
	}
	carOutput, err := s.Repository.GetCarByIdRepository(id)

	if err != nil {
		return domain.Car{}, err
	}
	return carOutput, nil
}
func (s *CarService) PutCarByIdService(car domain.Car) (domain.Car, error) {
	if car.Model == "" || car.Brand == "" || car.ID < 0 || car.OwnersCount < 0 || car.Mileage < 0 {
		return domain.Car{}, domain.ErrBadParamInput
	}
	carOutput, err := s.Repository.PutCarByIdRepository(car)

	if err != nil {
		return domain.Car{}, err
	}
	return carOutput, nil
}
func (s *CarService) PatchCarByIdService(fieldsToUpdate map[string]interface{}) (domain.Car, error) {
	if model, ok := fieldsToUpdate["model"]; ok && model.(string) == "" {
		return domain.Car{}, domain.ErrBadParamInput
	}

	if brand, ok := fieldsToUpdate["brand"]; ok && brand.(string) == "" {
		return domain.Car{}, domain.ErrBadParamInput
	}

	if mileage, ok := fieldsToUpdate["mileage"]; ok && mileage.(float64) < 0 {
		return domain.Car{}, domain.ErrBadParamInput
	}

	if ownersCount, ok := fieldsToUpdate["owners_count"]; ok && ownersCount.(float64) < 0 {
		return domain.Car{}, domain.ErrBadParamInput
	}

	carOutput, err := s.Repository.PatchCarByIdRepository(fieldsToUpdate)

	if err != nil {
		return domain.Car{}, err
	}
	return carOutput, nil

}
func (s *CarService) DeleteCarService(id int) error {
	if err := s.Repository.DeleteCarRepository(id); err != nil {
		return err
	}

	return nil
}
