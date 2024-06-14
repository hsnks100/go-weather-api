package ultrasrtfcst

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

// 초단기예보정보를 조회하기 위해 발표일자, 발표시각, 예보지점 X 좌표, 예보지점 Y 좌표의 조회 조건으로 자료구분코드, 예보값, 발표일자, 발표시각, 예보지점 X 좌표, 예보지점 Y 좌표의 정보를 조회하는 기능
func WeahterInfo(serviceKey, address string) ([]weather.WeatherInfo, error) {
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
	apiURL := "http://apis.data.go.kr/1360000/VilageFcstInfoService_2.0/getUltraSrtFcst?" + params.Encode()
	items, err := getWeatherData(apiURL)
	if err != nil {
		return nil, fmt.Errorf("API 요청 중 에러 발생: %v", err)
	}
	ret := query.WeatherDataToWeatherInfo(items)
	f := util.KoreaNow().Add(16 * time.Hour)
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
