package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hydde7/goexpert-challenge-2/service-b/internal/cfg"
	"github.com/hydde7/goexpert-challenge-2/service-b/internal/models"
)

func FreeWeatherRequest(ctx context.Context, city string) (*models.FreeWeatherPayload, error) {
	city = strings.ReplaceAll(strings.ToLower(city), " ", "+")
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", cfg.FreeWeather.ApiKey, city)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na requisição: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result models.FreeWeatherPayload
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
