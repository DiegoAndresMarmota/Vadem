package main

import (
	"context"
	"reflect"
	"testing"
)

func TestNewLogin(t *testing.T) {
	type args struct {
		entry VademService
	}
	tests := []struct {
		name string
		args args
		want VademService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogin(tt.args.entry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogin_GetVademecum(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name           string
		lo             *Login
		args           args
		wantMedicament *Medicament
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMedicament, err := tt.lo.GetVademecum(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login.GetVademecum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMedicament, tt.wantMedicament) {
				t.Errorf("Login.GetVademecum() = %v, want %v", gotMedicament, tt.wantMedicament)
			}
		})
	}
}
