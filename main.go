package main

import (
	"fmt"

	"example.com/filemanager"
	"example.com/prices"
)


func main() {

	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxrate := range taxRates{
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxrate * 100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxrate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job.")
		}

	}



}
