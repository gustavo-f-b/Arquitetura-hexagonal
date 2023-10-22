package application_test

import (
	"testing"

	"github.com/gustavo-f-b/Arquitetura-hexagonal/src/application"
	uuid "github.com/satori/go.uuid"
)

func TestProduct_Enable(t *testing.T) {
	tests := []struct {
		name    string
		product *application.Product
		wantErr bool
	}{
		{
			name:    "Valid Product",
			product: &application.Product{Name: "Widget", Price: 10.0},
			wantErr: false,
		},
		{
			name:    "Invalid Product",
			product: &application.Product{Name: "Widget", Price: -5.0},
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
		product *application.Product
		wantErr bool
	}{
		{
			name:    "Invalid Product",
			product: &application.Product{Name: "Widget", Price: 10.0},
			wantErr: true,
		},
		{
			name:    "Invalid Product",
			product: &application.Product{Name: "Widget", Price: -5.0},
			wantErr: true,
		},
		{
			name:    "Valid Product",
			product: &application.Product{Name: "Widget", Price: 0},
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

func TestProduct_IsValid(t *testing.T) {
	tests := []struct {
		name    string
		product *application.Product
		want    bool
		wantErr bool
	}{
		{
			name:    "Invalid Product",
			product: &application.Product{Id: uuid.NewV4().String(), Name: "Widget", Status: application.DISABLED, Price: -10.0},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.product.IsValid()
			if (err != nil) != tt.wantErr {
				t.Errorf("Product.IsValid() error = %v, wantErr %v", err, tt.wantErr)
				return

			}
			if got != tt.want {
				t.Errorf("Product.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
