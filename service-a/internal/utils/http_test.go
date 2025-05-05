package utils

import (
	"context"
	"net/http"
	"testing"
)

func TestGetJSON(t *testing.T) {
	type args struct {
		ctx    context.Context
		url    string
		result interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with valid URL",
			args: args{
				ctx: context.Background(),
				url: "https://api.example.com/data",
				result: &struct {
					Data string `json:"data"`
				}{},
			},
			wantErr: false,
		},
		{
			name: "Test with invalid URL",
			args: args{
				ctx: context.Background(),
				url: "invalid-url",
				result: &struct {
					Data string `json:"data"`
				}{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetJSON(tt.args.ctx, tt.args.url, tt.args.result)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp.StatusCode != http.StatusOK && !tt.wantErr {
				t.Errorf("GetJSON() gotResp = %v, want %v", gotResp.StatusCode, http.StatusOK)
			}
		})
	}
}
