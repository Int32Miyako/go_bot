package gis_meteo

import (
	"encoding/json"
	"io/ioutil"
	"meteo_bot/internal/entities"
	"net/http"
	"net/url"
	"strconv"
)

type GisMeteoAPI struct {
	client      *http.Client
	coordinates *entities.Coordinates
}

func NewGisMeteoAPI() *GisMeteoAPI {
	return &GisMeteoAPI{
		client: &http.Client{},
	}
}

func (_ *GisMeteoAPI) ServeGisMeteo(latitude, longitude string) (string, error) {
	query := "https://api.gismeteo.net/v2/weather/current/"
	params := url.Values{}
	params.Add("latitude", latitude)
	params.Add("longitude", longitude)
	params.Add("lang", "ru")
	query += "?" + params.Encode()

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

	data := &entities.GisResponse{}
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return "", err
	}

	return strconv.FormatFloat(data.Response.Temperature.Air.C, 'f', 6, 64), nil
}
