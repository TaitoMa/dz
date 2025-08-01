package translate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Result struct {
	Translation string `json:"translation"`
}

func GetTranslate(text string) (string, error) {
	url1 := fmt.Sprintf(
		"https://lingva.ml/api/v1/en/ru/%s",
		url.QueryEscape(text),
	)

	resp, err := http.Get(url1)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result Result
	json.Unmarshal(body, &result)

	return result.Translation, nil
}
