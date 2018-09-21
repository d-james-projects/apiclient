package main

import (
	"fmt"
	"log"
)

func main() {
	result, err := GetFuelMix()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s -> %s\n", result.Data.From, result.Data.To)
	for _, item := range result.Data.GenMix {
		fmt.Printf("%s %.1f\n",
			item.Fuel, item.Perc)
	}
}
