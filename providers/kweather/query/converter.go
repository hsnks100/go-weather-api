package query

import (
	"fmt"
	"sort"
	"time"

	"github.com/hsnks100/go-weather-api/providers/kweather/data"
	"github.com/hsnks100/go-weather-api/weather"
)

// 문자열을 날짜와 시간으로 파싱하여 time.Time으로 반환
func parseDateTime(dateStr, timeStr string) (time.Time, error) {
	layout := "20060102.1500" // 날짜와 시간 형식
	dtStr := dateStr + "." + timeStr
	// KST 시간대로 파싱하기 위해 Asia/Seoul 시간대 설정
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(layout, dtStr, loc)
}

// 문자열을 정수로 파싱
func parseInt(s string) int {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	if err != nil {
	}
	return result
}
func WeatherDataToWeatherInfo(items []data.WeatherResponseItem) []weather.WeatherInfo {
	timeTable := map[string]weather.WeatherInfo{}
	for _, item := range items {
		key := fmt.Sprintf("%s.%s", item.FcstDate, item.FcstTime)
		v := timeTable[key]
		v.FcstDateTime, _ = parseDateTime(item.FcstDate, item.FcstTime)
		switch item.Category {
		case "POP":
			v.Pop = parseInt(item.FcstValue)
		case "PTY":
			v.PrecipType = item.FcstValue
		case "PCP":
			v.Rainfall = parseInt(item.FcstValue)
		case "RN1":
			v.Rainfall = parseInt(item.FcstValue)
		case "REH":
			v.Humidity = parseInt(item.FcstValue)
		case "SNO":
			v.Snowfall = parseInt(item.FcstValue)
		case "SKY":
			v.Sky = item.FcstValue
		case "TMP":
			v.Temperature = parseInt(item.FcstValue)
		case "T1H":
			v.Temperature = parseInt(item.FcstValue)
		case "TMN":
			v.MinTemp = parseInt(item.FcstValue)
		case "TMX":
			v.MaxTemp = parseInt(item.FcstValue)
		case "UUU":
			v.WindEast = parseInt(item.FcstValue)
		case "VVV":
			v.WindNorth = parseInt(item.FcstValue)
		case "WAV":
			v.Wave = parseInt(item.FcstValue)
		case "VEC":
			v.WindDir = parseInt(item.FcstValue)
		case "WSD":
			v.WindSpeed = parseInt(item.FcstValue)
		case "LGT":
			v.Lightning = parseInt(item.FcstValue)
		default:
			fmt.Println("!!! Unknown category:", item.Category)
		}
		timeTable[key] = v
	}

	ret := make([]weather.WeatherInfo, 0, len(timeTable))
	for _, v := range timeTable {
		ret = append(ret, v)
	}
	// sort by time
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].FcstDateTime.Before(ret[j].FcstDateTime)
	})
	return ret
}
