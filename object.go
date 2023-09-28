package main

type Medicament struct {
	Name       string `json:"name"`
	Country    string `json:"country"`
	Laboratory string `json:"laboratory"`
	Suspension string `json:"suspension"`
	Form       string `json:"form"`
	ATC        string `json:"atc"`
}

type Meds struct {
	Meds []Medicament
}
