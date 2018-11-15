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

type GetRespRegional struct {
	Datatop []*ContentRegion `json:"datatop"`
}

type GetResp24Regional struct {
	Datatop Content24Region `json:"datatop"`
}

type Content24Region struct {
	RegionID  int          `json:"regionid"`
	Shortname string       `json:"shortname"`
	Data      []*ContentUK `json:"data"`
}

type Forecast struct {
	Fcast string `json:"forecast"`
	Idx   string `json:"index"`
}

type ContentUK struct {
	From   string     `json:"from"`
	To     string     `json:"to"`
	GenMix []*FuelMix `json:"generationmix"`
}

type ContentRegion struct {
	RegionID  int          `json:"regionid"`
	Shortname string       `json:"shortname"`
	Data      []*ContentUK `json:"data"`
}

type FuelMix struct {
	Fuel string  `json:"fuel"`
	Perc float32 `json:"perc"`
}
