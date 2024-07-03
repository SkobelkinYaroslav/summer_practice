package service

import "summer_practice/internal/domain"

type CarRepository interface {
	CreateCarRepository(car domain.Car) (domain.Car, error)
	GetAllCarsRepository() ([]domain.Car, error)
	GetCarByIdRepository(id int) (domain.Car, error)
	UpdateCarByIdRepository(car domain.Car) (domain.Car, error)
	PatchCarByIdRepository(car domain.Car) (domain.Car, error)
	DeleteCarRepository(id int) error
}

type CarService struct {
	Repository CarRepository
}

func NewCarService(repo CarRepository) *CarService {
	return &CarService{
		Repository: repo,
	}
}

func (s *CarService) CreateCarService(car domain.Car) (domain.Car, error) {
	carOutput, err := s.Repository.CreateCarRepository(car)
	if err != nil {
		return domain.Car{}, err
	}
	return carOutput, nil
}
func (s *CarService) GetAllCarsService() ([]domain.Car, error) {
	carArray, err := s.Repository.GetAllCarsRepository()
	if err != nil {
		return nil, err
	}
	return carArray, nil
}
func (s *CarService) GetCarByIdService(id int) (domain.Car, error) {
	carOutput, err := s.Repository.GetCarByIdRepository(id)

	if err != nil {
		return domain.Car{}, err
	}
	return carOutput, nil
}
func (s *CarService) UpdateCarByIdService(car domain.Car) (domain.Car, error) {
	carOutput, err := s.Repository.UpdateCarByIdRepository(car)

	if err != nil {
		return domain.Car{}, err
	}
	return carOutput, nil
}
func (s *CarService) PatchCarByIdService(car domain.Car) (domain.Car, error) {
	carOutput, err := s.Repository.PatchCarByIdRepository(car)

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
