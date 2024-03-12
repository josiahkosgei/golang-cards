package models

type Card struct {
	Id          string `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Status      string `json:"Status"`
	CreatedDate string `json:"CreatedDate"`
}
