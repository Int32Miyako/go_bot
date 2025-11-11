package open_meteo

import "net/http"

//https://open-meteo.com/en/docs

type OpenMeteoAPI struct {
	client *http.Client
	url    string
}

func NewOpenMeteoAPI() *OpenMeteoAPI {
	return &OpenMeteoAPI{
		client: &http.Client{},
	}
}

func (o *OpenMeteoAPI) GetWeather(latitude, longitude string) {
	url := "https://api.open-meteo.com/v1/forecast"
	body := "?latitude=" + latitude + "&longitude=" + longitude
	url += body
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	resp, err := o.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

}
