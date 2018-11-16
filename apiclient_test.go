package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var (
	testclient *Client
	testmux    *http.ServeMux
	testserver *httptest.Server
)

func setup() {
	testclient = NewClient()
	testmux = http.NewServeMux()
	testserver = httptest.NewServer(testmux)
	url, _ := url.Parse(testserver.URL)
	testclient.BaseURL = url
}

func teardown() {
	testserver.Close()
}
func TestRealGetCurrentUK(t *testing.T) {
	apiClient := NewClient()

	_, err := apiClient.GetCurrentUK(context.Background())
	if err != nil {
		t.Errorf("Api call error %v", err)
	}
}

func TestNewClient(t *testing.T) {
	apiClient := NewClient()

	if apiClient.BaseURL.String() != CarbonIntensityBaseURL {
		t.Errorf("NewClient BaseURL = %v, want %v", apiClient.BaseURL.String(), CarbonIntensityBaseURL)
	}

	if apiClient.UserAgent != userAgent {
		t.Errorf("NewClient UserAgent = %v, want %v", apiClient.UserAgent, userAgent)
	}
}

func TestGetCurrentUK(t *testing.T) {
	setup()
	defer teardown()

	testmux.HandleFunc("/generation", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"data":{"from":"fromdate","to":"todate",
			  "generationmix":[{"fuel":"biomass","perc":1},
				{"fuel":"coal","perc":2},
				{"fuel": "imports","perc":3}]}}`)
	})

	currentUK, err := testclient.GetCurrentUK(context.Background())
	if err != nil {
		t.Errorf("GetCurrentUK returned error: %v", err)
	}

	want := GetRespUK{}
	wantfuel := []FuelMix{{Fuel: "biomass", Perc: 1}, {Fuel: "coal", Perc: 2}, {Fuel: "imports", Perc: 3}}
	want.Data.From = "fromdate"
	want.Data.To = "todate"

	if strings.Compare(currentUK.Data.From, want.Data.From) != 0 {
		t.Errorf("TestGetCurrentUK From = %v, want %v", currentUK.Data.From, want.Data.From)
	}

	if strings.Compare(currentUK.Data.To, want.Data.To) != 0 {
		t.Errorf("TestGetCurrentUK To = %v, want %v", currentUK.Data.To, want.Data.To)
	}

	for i := 0; i < 3; i++ {
		if strings.Compare(currentUK.Data.GenMix[i].Fuel, wantfuel[i].Fuel) != 0 {
			t.Errorf("TestGetCurrentUK Fuel = %v, want %v", currentUK.Data.GenMix[i].Fuel, wantfuel[i].Fuel)
		}
		if currentUK.Data.GenMix[i].Perc != wantfuel[i].Perc {
			t.Errorf("TestGetCurrentUK Fuel = %v, want %v", currentUK.Data.GenMix[i].Perc, wantfuel[i].Perc)
		}
	}
}
