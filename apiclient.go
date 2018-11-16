package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func NewClient() *Client {
	doer := http.DefaultClient
	baseURL, err := url.Parse(CarbonIntensityBaseURL)
	if err != nil {
		log.Fatal(err)
	}

	client := &Client{httpClient: doer, BaseURL: baseURL, UserAgent: userAgent}
	return client
}

func (c *Client) GetCurrentUK(ctx context.Context) (*GetRespUK, error) {
	req, err := c.newRequest(ctx, "GET", "/generation", nil)
	if err != nil {
		return nil, err
	}

	var getUKFuelMixResult GetRespUK
	_, err = c.do(req, &getUKFuelMixResult)
	return &getUKFuelMixResult, err
}

func (c *Client) GetCurrentRegion(ctx context.Context, regionid string) (*GetRespRegional, error) {
	req, err := c.newRequest(ctx, "GET",
		"/regional/regionid/"+strings.Trim(regionid, " "),
		nil)
	if err != nil {
		return nil, err
	}

	var getRegionalFuelMixResult GetRespRegional
	_, err = c.do(req, &getRegionalFuelMixResult)
	return &getRegionalFuelMixResult, err
}

func (c *Client) GetRegion24(ctx context.Context, regionid string, start24 string) (*GetResp24Regional, error) {
	req, err := c.newRequest(ctx, "GET",
		"/regional/intensity/"+strings.Trim(start24, " ")+
			"/pt24h/regionid/"+strings.Trim(regionid, " "),
		nil)
	if err != nil {
		return nil, err
	}

	var getRegional24FuelMixResult GetResp24Regional
	_, err = c.do(req, &getRegional24FuelMixResult)
	return &getRegional24FuelMixResult, err
}

func (c *Client) newRequest(ctx context.Context, method string, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	fmt.Printf("Performing %v on URL %v\n", method, u.String())

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// need to modify json
	buf := new(bytes.Buffer)
	tee := io.TeeReader(resp.Body, buf)
	copyBuf := new(bytes.Buffer)
	copyBuf.ReadFrom(tee)
	respStr := copyBuf.String()

	var newStr string
	if strings.Count(respStr, KeyReplace) > 1 {
		newStr = strings.Replace(respStr, KeyReplace, KeyNew, 1)
	} else {
		newStr = respStr
	}

	tee = bytes.NewReader([]byte(newStr))

	err = json.NewDecoder(tee).Decode(v)
	return resp, err
}
