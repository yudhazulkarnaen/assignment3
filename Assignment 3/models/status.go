package models

type Status struct {
	Status StatusBody `json:"status"`
}

type StatusBody struct {
	Water int
	Wind  int
}
