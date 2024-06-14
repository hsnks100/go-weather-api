package weather

import "time"

// WeatherInfo holds detailed current weather information.
type WeatherInfo struct {
	Pop          int       // 강수확률 (Precipitation Probability)
	PrecipType   string    // 강수형태 (Precipitation Type)
	Rainfall     int       // 1시간 강수량 (Rainfall in mm per hour)
	Humidity     int       // 습도 (Humidity percentage)
	Snowfall     int       // 1시간 신적설 (Snowfall in cm per hour)
	Sky          string    // 하늘상태 (Sky condition)
	Temperature  int       // 1시간 기온 (Temperature in Celsius)
	MinTemp      int       // 일 최저기온 (Minimum daily temperature)
	MaxTemp      int       // 일 최고기온 (Maximum daily temperature)
	WindEast     int       // 동서풍속 (East-West wind speed in km/h)
	WindNorth    int       // 남북풍속 (North-South wind speed in km/h)
	Wave         int       // 파고 (Wave height in meters)
	WindDir      int       // 풍향 (Wind direction in degrees)
	WindSpeed    int       // 풍속 (Wind speed in km/h)
	Lightning    int       // 번개 (Lightning) kA
	FcstDateTime time.Time // 예보 일시 (Forecast datetime)
}
