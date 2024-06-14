package vilagefcst

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/hsnks100/go-weather-api/providers/kweather/data"
)

const (
	serviceKey = "**********"
)

func TestName(t *testing.T) {
}

// printCoordinates 함수는 Coordinates 타입을 문자열로 포맷팅합니다.
func printCoordinates(c data.Coordinates) string {
	return fmt.Sprintf("{%d, %d}", c.X, c.Y)
}

// printRegionMap 함수는 RegionMap 타입을 문자열로 포맷팅합니다.
func printRegionMap(rm data.RegionMap) string {
	keys := make([]string, 0, len(rm))
	for key := range rm {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	result := "RegionMap{\n"
	for _, key := range keys {
		result += fmt.Sprintf("\t\"%s\": %s,\n", key, printCoordinates(rm[key]))
	}
	result += "}"
	return result
}

// printCityMap 함수는 CityMap 타입을 문자열로 포맷팅합니다.
func printCityMap(cm data.CityMap) string {
	keys := make([]string, 0, len(cm))
	for key := range cm {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	result := "CityMap{\n"
	for _, key := range keys {
		result += fmt.Sprintf("\t\"%s\": %s,\n", key, printRegionMap(cm[key]))
	}
	result += "}"
	return result
}

// printProvinceMap 함수는 ProvinceMap 타입을 문자열로 포맷팅합니다.
func printProvinceMap(pm data.ProvinceMap) string {
	keys := make([]string, 0, len(pm))
	for key := range pm {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	result := "ProvinceMap{\n"
	for _, key := range keys {
		result += fmt.Sprintf("\t\"%s\": %s,\n", key, printCityMap(pm[key]))
	}
	result += "}"
	return result
}

func TestGenerateWeaterZone(t *testing.T) {
	in, err := os.ReadFile("in.txt")
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(in)))
	provinceMap := data.ProvinceMap{}

	for scanner.Scan() {
		input := scanner.Text()
		if input == "end" {
			break
		}

		parts := strings.Fields(input)
		if len(parts) < 3 {
			fmt.Println("잘못된 형식입니다. 최소 도시와 좌표가 필요합니다. 다시 입력해주세요.")
			continue
		}

		province := parts[0]
		city := ""
		region := ""
		x := ""
		y := ""

		// 입력 파싱
		switch len(parts) {
		case 3:
			x = parts[1]
			y = parts[2]
		case 4:
			city = parts[1]
			x = parts[2]
			y = parts[3]
		case 5:
			city = parts[1]
			region = parts[2]
			x = parts[3]
			y = parts[4]
		default:
			fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
			continue
		}

		// 좌표 값 변환
		xCoord, errX := strconv.Atoi(x)
		yCoord, errY := strconv.Atoi(y)
		if errX != nil || errY != nil {
			fmt.Println("좌표는 숫자여야 합니다.")
			continue
		}

		coord := data.Coordinates{X: xCoord, Y: yCoord}

		// 도시 맵 생성
		if _, ok := provinceMap[province]; !ok {
			provinceMap[province] = data.CityMap{}
		}

		// 구 맵 생성
		if _, ok := provinceMap[province][city]; !ok {
			provinceMap[province][city] = data.RegionMap{}
		}

		// 좌표 설정
		provinceMap[province][city][region] = coord
	}
	fmt.Println(printProvinceMap(provinceMap))
}

func TestFindCoordinates(t *testing.T) {
	coord, err := data.FindCoordinates(data.Province, "서울특별시 금천구 가산동")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(coord)

	coord2, err := data.FindAnyMatchingCoordinates(data.Province, "논현1동")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(coord2)

	coord3, err := data.FindAnyMatchingCoordinates(data.Province, "가산동")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(coord3)
}
func TestWeather(t *testing.T) {
	we, err := WeahterInfo(serviceKey, "가산동")
	if err != nil {
		t.Fatal(err)
	}
	for _, w := range we {
		fmt.Printf("%+v\n", w)
	}
}
