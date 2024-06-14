package kweather

import (
	"fmt"
	"time"

	"github.com/hsnks100/go-weather-api/providers/kweather/query/vilagefcst"
	"github.com/hsnks100/go-weather-api/weather"
)

type Client struct {
	APIKey string
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func (c *Client) GetCurrentWeather(location string) (*weather.WeatherInfo, error) {
	info, err := vilagefcst.WeahterInfo(c.APIKey, location, 16*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("weater info error: %v", err)
	}
	if len(info) == 0 {
		return nil, fmt.Errorf("no weather info found")
	}
	return &info[0], nil
}

func (c *Client) GetDetailedForecast(location string, duration time.Duration) ([]weather.WeatherInfo, error) {
	info, err := vilagefcst.WeahterInfo(c.APIKey, location, duration)
	if err != nil {
		return nil, fmt.Errorf("weater info error: %v", err)
	}
	if len(info) == 0 {
		return nil, fmt.Errorf("no weather info found")
	}
	return info, nil
}

func (c *Client) ReportWeather(w []weather.WeatherInfo) string {
	if len(w) == 0 {
		return "날씨 정보가 없습니다."
	}
	maxTemp := -100
	minTemp := 100
	maxHumidity := 0
	rainFall := 0
	maxPOP := 0
	skyMap := make(map[string]int)
	for _, v := range w {
		// 맑음(1), 구름많음(3), 흐림(4)
		switch v.Sky {
		case "1":
			skyMap["맑음"]++
		case "3":
			skyMap["구름많음"]++
		case "4":
			skyMap["흐림"]++
		}
		if v.Temperature > maxTemp {
			maxTemp = v.Temperature
		}
		if v.Temperature < minTemp {
			minTemp = v.Temperature
		}
		if v.Humidity > maxHumidity {
			maxHumidity = v.Humidity
		}
		if v.Rainfall > rainFall {
			rainFall = v.Rainfall
		}
		if v.Pop > maxPOP {
			maxPOP = v.Pop
		}
	}
	skyType := ""
	maxSky := 0
	for k, v := range skyMap {
		if v > maxSky {
			maxSky = v
			skyType = k
		}
	}

	format := `[날씨 요약] 
날짜 및 시간: %v~%v
기온: %d~%d°C
강수확률: %d
강수량: %dmm
하늘상태: %s
습도: %d%%`
	return fmt.Sprintf(format, w[0].FcstDateTime.Format("2006-01-02 15:04:05.999999999"), w[len(w)-1].FcstDateTime.Format("2006-01-02 15:04:05.999999999"), minTemp, maxTemp, maxPOP, rainFall, skyType, maxHumidity)
}
