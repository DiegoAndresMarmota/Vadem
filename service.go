package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type VademService interface {
	GetVademecum(context.Context) (*Medicament, error)
}


type VademIndex struct {
	url string

}

func NewVademIndex(url string) VademService{
	return &VademIndex{
		url: url,
	}
}


func (v *VademIndex) GetVademecum(ctx context.Context) (*Medicament, error) {
	resp, err := http.Get(v.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	medicament := &Medicament{
			Name: resp.Header.Get("name"),
			Country: resp.Header.Get("country"),
			Laboratory: resp.Header.Get("laboratory"),
			Suspension: resp.Header.Get("suspension"),
			Form: resp.Header.Get("form"),
			ATC: resp.Header.Get("ATC"),
	}
	if err := json.NewDecoder(resp.Body).Decode(&medicament); err != nil {
		return nil, err
	}

	return medicament, nil

}