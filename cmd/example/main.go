package main

import (
	"fmt"
	"time"

	"github.com/hsnks100/go-weather-api/providers/kweather"
)

func main() {
	const serviceKey = "**********"
	ks := kweather.NewClient(serviceKey)
	wi2, err := ks.GetDetailedForecast("가산동", 16*time.Hour)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range wi2 {
		fmt.Println(v)
	}
	fmt.Println("#2: ", ks.ReportWeather(wi2))
}
