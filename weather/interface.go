package weather

import "time"

// WeatherProvider defines the interface for fetching detailed weather data.
type WeatherProvider interface {
	GetCurrentWeather(location string) (*WeatherInfo, error)                            // 현재 날씨 정보를 가져옵니다.
	GetDetailedForecast(location string, duration time.Duration) ([]WeatherInfo, error) // 지정된 위치의 상세한 날씨 예보를 배열 형태로 가져옵니다.
	ReportWeather(info []WeatherInfo) string                                            // 날씨 정보를 보고서 형태로 반환합니다.
}
