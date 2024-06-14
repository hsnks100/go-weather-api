package util

import (
	"fmt"
	"time"
)

func KoreaNow() time.Time {
	utcTime := time.Now().UTC()

	// Asia/Seoul 시간대로 변경
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		fmt.Println("Error:", err)
		return time.Time{}
	}
	// KST 시간으로 변환
	kstTime := utcTime.In(location)
	return kstTime
}
func ClosestDateTime() (string, string) {
	currentTime := KoreaNow()
	// Base_time 리스트
	baseTimes := []string{"0200", "0500", "0800", "1100", "1400", "1700", "2000", "2300"}

	// 현재 시간과의 차이 계산
	var closestTime string

	for _, baseTime := range baseTimes {
		// 현재 날짜와 시간을 가져와서 시간으로 변환
		currentYear, currentMonth, currentDay := currentTime.Date()
		baseHour, _ := time.Parse("1504", baseTime)
		baseDateTime := time.Date(currentYear, currentMonth, currentDay, baseHour.Hour(), 0, 0, 0, currentTime.Location())
		// 현재 시간과의 차이 계산
		diff := baseDateTime.Sub(currentTime)
		if diff < 0 {
			closestTime = baseTime
		}
	}
	var formattedTime string
	if closestTime == "" {
		closestTime = "2300"
		formattedTime = currentTime.Add(-24 * time.Hour).Format("20060102")
	} else {
		formattedTime = currentTime.Format("20060102")
	}
	return formattedTime, closestTime
}
