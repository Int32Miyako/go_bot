package open_meteo

import "net/http"

//https://open-meteo.com/en/docs

type OpenMeteoAPI struct {
	client *http.Client
}

func NewOpenMeteoAPI() *OpenMeteoAPI {
	return &OpenMeteoAPI{
		client: &http.Client{},
	}
}
