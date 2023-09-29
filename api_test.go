package main

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewApiVademecum(t *testing.T) {
	type args struct {
		explorer VademService
	}
	tests := []struct {
		name string
		args args
		want *ApiVademecum
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewApiVademecum(tt.args.explorer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewApiVademecum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApiVademecum_handleGetVademecum(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		api  *ApiVademecum
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.api.handleGetVademecum(tt.args.w, tt.args.r)
		})
	}
}

func Test_allMedicamentsJSON(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		api int
		v   any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := allMedicamentsJSON(tt.args.w, tt.args.api, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("allMedicamentsJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApiVademecum_Iniciate(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		api     *ApiVademecum
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.api.Iniciate(tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("ApiVademecum.Iniciate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
