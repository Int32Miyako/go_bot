package gis_meteo

import "net/http"

type GisMeteoAPI struct {
	client *http.Client
}
