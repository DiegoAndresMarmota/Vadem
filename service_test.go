package main

import (
	"context"
	"reflect"
	"testing"
)

func TestNewVademIndex(t *testing.T) {
	type args struct {
		url string
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
			if got := NewVademIndex(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVademIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVademIndex_GetVademecum(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		v       *VademIndex
		args    args
		want    *Medicament
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.GetVademecum(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("VademIndex.GetVademecum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VademIndex.GetVademecum() = %v, want %v", got, tt.want)
			}
		})
	}
}
