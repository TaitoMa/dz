package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"
)

func main() {
	city := flag.String("city", "", "Город")
	format := flag.Int("format", 1, "Формат")
	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(weather.GetWeather(*geoData, *format))
}
