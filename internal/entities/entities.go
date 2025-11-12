package entities

type Coordinates struct {
	Latitude  string // Широта
	Longitude string // Долгота
}

type GisResponse struct {
	Response struct {
		Temperature struct {
			Air struct {
				C float64 `json:"C"`
			} `json:"air"`
		} `json:"temperature"`
		Description struct {
			Full string `json:"full"`
		} `json:"description"`
	} `json:"response"`
}

type OpenWeatherResponse struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Current        struct {
		Dt         int64   `json:"dt"`
		Sunrise    int64   `json:"sunrise"`
		Sunset     int64   `json:"sunset"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Uvi        float64 `json:"uvi"`
		Clouds     int     `json:"clouds"`
		Visibility int     `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindDeg    int     `json:"wind_deg"`
		WindGust   float64 `json:"wind_gust"`
		Weather    []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"current"`
}
