package utils

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestValidateCEP(t *testing.T) {
	type args struct {
		cep     string
		isValid bool
	}

	tests := []args{
		{"12345678", true},
		{"12345-1234", false},
		{"1234123", false},
		{"123451", false},
		{"123d5-123", false},
		{"1234512d", false},
		{"19000000", true},
		{"29000000", true},
	}

	for _, tt := range tests {
		isValid := ValidateCEP(tt.cep)
		assert.Equal(t, tt.isValid, isValid)
	}

}
