package vilagefcst

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/hsnks100/go-weather-api/providers/kweather/data"
	"github.com/hsnks100/go-weather-api/providers/kweather/query"
	"github.com/hsnks100/go-weather-api/providers/kweather/util"
	"github.com/hsnks100/go-weather-api/weather"
)

func ReportWeather(w []weather.WeatherInfo) string {
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

	return fmt.Sprintf(`[날씨 요약] 
날짜 및 시간: %v
기온: %d~%d°C
강수확률: %d
강수량: %dmm
하늘상태: %s
습도: %d%%`, w[0].FcstDateTime.Format("2006-01-02 15:04:05.999999999"), minTemp, maxTemp, maxPOP, rainFall, skyType, maxHumidity)
}

func WeahterInfo(serviceKey, address string, duration time.Duration) ([]weather.WeatherInfo, error) {
	coord, err := data.FindAnyMatchingCoordinates(data.Province, address)
	if err != nil {
		return nil, fmt.Errorf("failed to find coordinates: %v", err)
	}
	// API 요청에 필요한 매개변수 설정
	params := url.Values{}
	params.Set("serviceKey", serviceKey)
	params.Set("numOfRows", "500")
	params.Set("pageNo", "1")
	date, t := util.ClosestDateTime()
	params.Set("base_date", date)
	params.Set("base_time", t)
	// 58, 125: 가산동
	params.Set("nx", strconv.Itoa(coord.X))
	params.Set("ny", strconv.Itoa(coord.Y))
	params.Set("dataType", "json")

	// API 요청 보내기
	apiURL := "http://apis.data.go.kr/1360000/VilageFcstInfoService_2.0/getVilageFcst?" + params.Encode()
	items, err := getWeatherData(apiURL)
	if err != nil {
		return nil, fmt.Errorf("API 요청 중 에러 발생: %v", err)
	}
	ret := query.WeatherDataToWeatherInfo(items)
	f := util.KoreaNow().Add(duration)
	for i, v := range ret {
		if f.Before(v.FcstDateTime) {
			ret = ret[:i]
			break
		}
	}
	return ret, nil
}

// getWeatherData 함수는 주어진 API 엔드포인트에서 날씨 데이터를 가져옵니다.
func getWeatherData(apiURL string) ([]data.WeatherResponseItem, error) {
	// API 요청 보내기
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 응답 바디 읽기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("응답 읽기 에러: %v", err)
	}
	// 응답을 WeatherResponse 구조체로 언마샬링
	var weatherResp data.WeatherResponse
	err = json.Unmarshal(body, &weatherResp)
	if err != nil {
		return nil, fmt.Errorf("JSON 언마샬링 에러: %v, URL: %v ", string(body), apiURL)
	}
	// API 응답에서 Item 리스트 반환
	return weatherResp.Response.Body.Items.Item, nil
}
