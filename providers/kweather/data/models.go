package data

type WeatherResponse struct {
	Response WeatherResponseBody `json:"response"`
}

type WeatherResponseBody struct {
	Header WeatherResponseHeader   `json:"header"`
	Body   WeatherResponseBodyData `json:"body"`
}

type WeatherResponseHeader struct {
	ResultCode string `json:"resultCode"`
	ResultMsg  string `json:"resultMsg"`
}

type WeatherResponseBodyData struct {
	DataType   string               `json:"dataType"`
	Items      WeatherResponseItems `json:"items"`
	PageNo     int                  `json:"pageNo"`
	NumOfRows  int                  `json:"numOfRows"`
	TotalCount int                  `json:"totalCount"`
}

type WeatherResponseItems struct {
	Item []WeatherResponseItem `json:"item"`
}

type WeatherResponseItem struct {
	BaseDate  string `json:"baseDate"`
	BaseTime  string `json:"baseTime"`
	Category  string `json:"category"`
	FcstDate  string `json:"fcstDate"`
	FcstTime  string `json:"fcstTime"`
	FcstValue string `json:"fcstValue"`
	Nx        int    `json:"nx"`
	Ny        int    `json:"ny"`
}
