// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

package main

import (
	"encoding/json"
	"net/http"
)

// SearchIssues queries the GitHub issue tracker.
func GetFuelMix() (*GetResp, error) {
	// resp, err := http.Get(CarbonIntensityURL)
	// if err != nil {
	// 	return nil, err
	// }

	//!-
	// For long-term stability, instead of http.Get, use the
	// variant below which adds an HTTP request header indicating
	// that only version 3 of the GitHub API is acceptable.
	//
	//   req, err := http.NewRequest("GET", IssuesURL+"?q="+q, nil)
	//   if err != nil {
	//       return nil, err
	//   }
	//   req.Header.Set(
	//       "Accept", "application/vnd.github.v3.text-match+json")
	//   resp, err := http.DefaultClient.Do(req)
	//!+

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	// if resp.StatusCode != http.StatusOK {
	// 	resp.Body.Close()
	// 	return nil, fmt.Errorf("search query failed: %s", resp.Status)
	// }

	// defer rs.Body.Close()

	// bodyBytes, err := ioutil.ReadAll(rs.Body)
	// if err != nil {
	//     panic(err)
	// }

	// bodyString := string(bodyBytes)

	client := http.Client{}
	request, err := http.NewRequest("GET", CarbonIntensityURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var getFuelMixResult GetResp
	if err := json.NewDecoder(resp.Body).Decode(&getFuelMixResult); err != nil {
		return nil, err
	}
	return &getFuelMixResult, nil
}

//!-
