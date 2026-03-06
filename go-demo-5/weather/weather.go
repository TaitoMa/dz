package weather

import (
	"demo/weather/geo"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func GetWeather(geo geo.GeoData, format int) string {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	params.Add("format", strconv.Itoa(format))
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
