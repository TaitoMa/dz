package geo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type GeoData struct {
	City string `json:"city"`
	//Region  string `json:"region"`
	//Country string `json:"country"`
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		checked := CheckCity(city)
		if !checked {
			panic("Такого города нет")
		}
		return &GeoData{City: city}, nil
	}
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", "https://ipapi.co/json", nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %v", err)
	}
	req.Header.Set("User-Agent", "curl/7.64.1") // Простой User-Agent

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе: %v", err)
	}
	defer resp.Body.Close() // Закрываем Body только после успешного запроса

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка HTTP: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа: %v", err)
	}

	var geo GeoData
	if err := json.Unmarshal(body, &geo); err != nil {
		return nil, fmt.Errorf("ошибка разбора JSON: %v", err)
	}

	if geo.Error {
		return nil, fmt.Errorf("API вернуло ошибку: %s", geo.Message)
	}

	return &geo, nil
}

type CityPopulationResponse struct {
	Error bool `json:"error"`
}

func CheckCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	resp, err := http.Post("http://ipapi.com/json/", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var response CityPopulationResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return false
	}
	return response.Error
}
