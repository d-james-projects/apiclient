// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 110.

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package main

const CarbonIntensityURL = "https://api.carbonintensity.org.uk/generation"

type GetResp struct {
	Data Content
}

type Content struct {
	From   string     `json:"from"`
	To     string     `json:"to"`
	GenMix []*FuelMix `json:"generationmix"`
}

type FuelMix struct {
	Fuel string  `json:"fuel"`
	Perc float32 `json:"perc"`
}
