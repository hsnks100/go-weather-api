package openweather

import (
	"time"

	"github.com/hsnks100/go-weather-api/weather"
)

type Client struct {
	APIKey string
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

func (c *Client) GetCurrentWeather(location string) (*weather.WeatherInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) GetDetailedForecast(location string, duration time.Duration) ([]weather.WeatherInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) ReportWeather(info []weather.WeatherInfo) string {
	//TODO implement me
	panic("implement me")
}
