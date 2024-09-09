package model

type Employee struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
}
