package prices

import (
	"errors"
	"fmt"

	"example.com/conversion"
	"example.com/iomanager"
)


type TaxIncludedPriceJob struct {
		IOManager iomanager.IOManager `json:"-"`
		TaxRate float64 `json:"tax_rate"`
		InputPrices []float64 `json:"input_prices"`
		TaxIncludedPrices map[string]string `json:"tax_included"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return errors.New("err")
	}
	priceArray, err := conversion.StringsToFloats(lines)

	if err != nil {
		return errors.New("err")
	}
	job.InputPrices = priceArray

	return nil
}


func (job *TaxIncludedPriceJob) Process() error{

	err := job.LoadData()

	if err != nil{
		return err
	}

	result := make(map[string]string)
		for _, price := range job.InputPrices {
			taxIncludedPrice := price * (1 + job.TaxRate)
			result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
		}

	job.TaxIncludedPrices = result

return	job.IOManager.WriteResult(job)

}


func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob{
		return &TaxIncludedPriceJob{
			IOManager: iom,
			TaxRate: taxRate,
		}
}

