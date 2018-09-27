package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// func main() {
// 	result, err := GetCurrentUK()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s -> %s\n", result.Data.From, result.Data.To)
// 	for _, item := range result.Data.GenMix {
// 		fmt.Printf("%s %.1f\n",
// 			item.Fuel, item.Perc)
// 	}
// }

func main() {
	u, err := url.Parse(CarbonIntensityBaseURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Interrogating URL at: ")
	fmt.Println(u)
	client := http.Client{}

	apiClient := &Client{
		BaseURL: u,
	}
	//	apiClient.BaseURL = u
	apiClient.httpClient = &client

	r1, err := apiClient.GetCurrentUK(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s -> %s\n", r1.Data.From, r1.Data.To)
	for _, item := range r1.Data.GenMix {
		fmt.Printf("%s %.1f\n",
			item.Fuel, item.Perc)
	}

	r2, err := apiClient.GetCurrentRegion(context.Background(), "11")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r2.Datatop[0].Shortname, r2.Datatop[0].RegionID)
	for _, item := range r2.Datatop[0].Data[0].GenMix {
		fmt.Printf("%s %.1f\n",
			item.Fuel, item.Perc)
	}

	r3, err := apiClient.GetCurrentRegion(context.Background(), "16")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r3.Datatop[0].Shortname, r3.Datatop[0].RegionID)
	for _, item := range r3.Datatop[0].Data[0].GenMix {
		fmt.Printf("%s %.1f\n",
			item.Fuel, item.Perc)
	}

}
