package data

import (
	"fmt"
	"strings"
)

// FindAnyMatchingCoordinates는 사용자 입력에 따라 ProvinceMap에서 첫 번째 일치하는 좌표를 찾는 함수입니다.
func FindAnyMatchingCoordinates(provinceMap ProvinceMap, input string) (Coordinates, error) {
	if input == "" {
		return Coordinates{}, fmt.Errorf("input must not be empty")
	}
	// 입력값을 소문자로 통일
	lowerInput := strings.ToLower(input)

	// 도시 순회
	for _, cityMap := range provinceMap {
		// 구 순회
		for _, regionMap := range cityMap {
			// 동 순회
			for regionName, coords := range regionMap {
				if strings.Contains(strings.ToLower(regionName), lowerInput) {
					return coords, nil
				}
			}
		}
	}
	return Coordinates{}, fmt.Errorf("no matching coordinates found for '%s'", input)
}

// FindCoordinates는 주어진 주소에 따라 ProvinceMap에서 가장 비슷한 좌표를 찾는 함수입니다.
func FindCoordinates(provinceMap ProvinceMap, address string) (Coordinates, error) {
	// 주소를 공백으로 구분하여 분리
	parts := strings.Split(address, " ")
	if len(parts) < 3 {
		return Coordinates{}, fmt.Errorf("address must include city, district, and region")
	}

	city, district, region := parts[0], parts[1], parts[2]

	// 도시 찾기
	cityMap, ok := provinceMap[city]
	if !ok {
		return Coordinates{}, fmt.Errorf("city '%s' not found", city)
	}

	// 구 찾기
	regionMap, ok := cityMap[district]
	if !ok {
		return Coordinates{}, fmt.Errorf("district '%s' not found", district)
	}

	// 동 찾기
	coords, ok := regionMap[region]
	if !ok {
		return Coordinates{}, fmt.Errorf("region '%s' not found", region)
	}

	return coords, nil
}
