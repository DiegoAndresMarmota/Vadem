package main

import (
	"context"
	"fmt"
	"time"
)

type Login struct {
	entry VademService
}

func NewLogin(entry VademService) VademService {
	return &Login{
		entry: entry,
	}
}

func (lo *Login) GetVademecum(ctx context.Context) (medicament *Medicament, err error) {

	defer func(closeLogin time.Time){
		fmt.Printf("medicament=%s err=%v time=%v\n", medicament.Name, err, time.Since(closeLogin))
	}(time.Now())

	return lo.entry.GetVademecum(ctx)

}