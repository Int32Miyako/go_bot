package adapters

import (
	"encoding/json"
	"io/ioutil"
	"meteo_bot/internal/entities"
	"net/http"
	"net/url"
	"strconv"
)

// https://openweathermap.org/api/one-call-3

type OpenWeatherAPI struct {
	client *http.Client
	apiKey string
}

func NewOpenWeatherAPI(apiKey string) *OpenWeatherAPI {
	return &OpenWeatherAPI{
		client: &http.Client{},
		apiKey: apiKey,
	}
}

func (o *OpenWeatherAPI) ServeOpenWeather(latitude, longitude string) (string, error) {
	query := "https://api.openweathermap.org/data/2.5/onecall"
	params := url.Values{}
	params.Add("lat", latitude)
	params.Add("lon", longitude)
	params.Add("units", "metric") // цельсии
	params.Add("exclude", "hourly,daily,alerts")
	params.Add("lang", "ru")
	params.Add("appid", o.apiKey)
	query += "?" + params.Encode()

	// причины 401
	// https://openweathermap.org/faq#error401

	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	data := &entities.OpenWeatherResponse{}
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return "", err
	}

	return strconv.FormatFloat(data.Current.Temp, 'f', 6, 64), nil
}
