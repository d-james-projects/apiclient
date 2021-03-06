package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	apiClient := NewClient()

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

	r4, err := apiClient.GetRegion24(context.Background(), "16", "2018-09-30T12:00Z")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("region name %s\n", r4.Datatop.Shortname)
	for _, halfHourSlot := range r4.Datatop.Data {
		fmt.Printf("from %s\n", halfHourSlot.From)
		for _, item := range halfHourSlot.GenMix {
			fmt.Printf("%s %.1f\n",
				item.Fuel, item.Perc)
		}
	}
}
