package main

import (
	"net/http"
	"net/url"
)

const CarbonIntensityBaseURL = "https://api.carbonintensity.org.uk"
const KeyReplace = "\"data\""
const KeyNew = "\"datatop\""

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	httpClient *http.Client
}

type GetRespUK struct {
	Data ContentUK
}

type ContentUK struct {
	From   string     `json:"from"`
	To     string     `json:"to"`
	GenMix []*FuelMix `json:"generationmix"`
}

type GetRespRegional struct {
	Datatop []*ContentRegion
}

type ContentRegion struct {
	RegionID  int          `json:"regionid"`
	Shortname string       `json:"shortname"`
	Data      []*ContentUK `json:"data"`
}

type FuelMix struct {
	Fuel string
	Perc float32
}
