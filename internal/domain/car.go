package domain

type Car struct {
	ID          int    `json:"id"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	Mileage     int    `json:"mileage"`
	OwnersCount int    `json:"owners_count"`
}
