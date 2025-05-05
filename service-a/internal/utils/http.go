package utils

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func GetJSON(ctx context.Context, url string, result interface{}) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	json.Unmarshal(data, result)
	return resp, nil
}
