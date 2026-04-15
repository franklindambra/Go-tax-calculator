package main

import (
	"fmt"

	"example.com/filemanager"
	"example.com/prices"
)


func main() {

	taxRates := []float64{0, 0.7, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxrate := range taxRates{
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxrate * 100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxrate)
		go priceJob.Process(doneChans[index], errorChans[index])
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <- doneChans[index]: //let val go into void
		fmt.Println("done")
		}
	}


	// for _, doneChan := range doneChans {
	// 	fmt.Println(<- doneChan)
	// }



}
