package entity

type Catalog struct {
    Id int `json:"id"`
	Brand string `json:"brand"`
    Model string `json:"model"`
    Year int `json:"year"`
    Fuel string `json:"fuel"`
    Price float64 `json:"price"`
}