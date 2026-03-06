package cat

import (
	"encoding/json"
	"io"
	"net/http"
)

type Cat struct {
	Fact string `json:"fact"`
}

func GetCatFact() *Cat {
	resp, err := http.Get("https://catfact.ninja/fact/")
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var cat Cat
	err = json.Unmarshal(body, &cat)
	if err != nil {
		panic(err)
	}

	return &cat
}
