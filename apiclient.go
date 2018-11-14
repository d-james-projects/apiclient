package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// func GetCurrentUK() (*GetResp, error) {
// 	client := http.Client{}
// 	request, err := http.NewRequest("GET", CarbonIntensityURL, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp, err := client.Do(request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	var getFuelMixResult GetResp
// 	if err := json.NewDecoder(resp.Body).Decode(&getFuelMixResult); err != nil {
// 		return nil, err
// 	}
// 	return &getFuelMixResult, nil
// }

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

func (c *Client) GetRegion24(ctx context.Context, regionid string, start24 string) (*GetRespRegional, error) {
	req, err := c.newRequest(ctx, "GET",
		"/regional/intensity/"+strings.Trim(start24, " ")+
			"/pt24h/regionid/"+strings.Trim(regionid, " "),
		nil)
	if err != nil {
		return nil, err
	}

	var getRegional24FuelMixResult GetRespRegional
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
	//fmt.Printf(respStr)
	//b, err := json.MarshalIndent(copyBuf, "", "  ")
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//os.Stdout.Write(b)

	// more than 1 data field then replace the 1st one with datatop
	var newStr string
	if strings.Count(respStr, KeyReplace) > 1 {
		newStr = strings.Replace(respStr, KeyReplace, KeyNew, 1)
	} else {
		newStr = respStr
	}
	//fmt.Println(newStr)

	tee = bytes.NewReader([]byte(newStr))
	//b, err := json.MarshalIndent(tee, "", "  ")
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//os.Stdout.Write(b)

	content, _ := ioutil.ReadAll(tee)
	fmt.Println(string(content))

	err = json.NewDecoder(tee).Decode(v)
	return resp, err
}
