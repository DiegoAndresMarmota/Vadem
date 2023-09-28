package main

import (
	"context"
	"encoding/json"
	"net/http"
)


type ApiVademecum struct {
	explorer VademService
}


func NewApiVademecum(explorer VademService) *ApiVademecum {
	return &ApiVademecum{
		explorer: explorer,
	}
}


func (api *ApiVademecum) handleGetVademecum(w http.ResponseWriter, r *http.Request) {

	medicament, err := api.explorer.GetVademecum(context.Background())
	if err != nil {
		allMedicamentsJSON(w, http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
		return
	}

	allMedicamentsJSON(w, http.StatusOK, medicament)

}


func allMedicamentsJSON(w http.ResponseWriter, api int, v any) error {

	w.WriteHeader(api)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)

}


func (api *ApiVademecum) Iniciate(query string) error {

	http.HandleFunc("/", api.handleGetVademecum)
	return http.ListenAndServe(query, nil)

}