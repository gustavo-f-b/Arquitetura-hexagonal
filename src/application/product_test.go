package application

import (
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	tests := []struct {
		name    string
		product *Product
		wantErr bool
	}{
		{
			name:    "Valid Product",
			product: &Product{Name: "Widget", Price: 10.0},
			wantErr: false,
		},
		{
			name:    "Invalid Product",
			product: &Product{Name: "Widget", Price: -5.0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.product.Enable(); (err != nil) != tt.wantErr {
				t.Errorf("Product.Enable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProduct_Disable(t *testing.T) {
	tests := []struct {
		name    string
		product *Product
		wantErr bool
	}{
		{
			name:    "Invalid Product",
			product: &Product{Name: "Widget", Price: 10.0},
			wantErr: true,
		},
		{
			name:    "Invalid Product",
			product: &Product{Name: "Widget", Price: -5.0},
			wantErr: true,
		},
		{
			name:    "Valid Product",
			product: &Product{Name: "Widget", Price: 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.product.Disable(); (err != nil) != tt.wantErr {
				t.Errorf("Product.Disable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
