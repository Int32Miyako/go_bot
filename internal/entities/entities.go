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
