package repo

import (
	"testing"
)

var (
	hasilTax    float32 = 30
	hasilAmount float32 = 1030
)

func TestTaxCalculation(t *testing.T) {

	items, _ := TaxCalculate(1)
	t.Logf("Volume : %.2f", items.Tax)

	if items.Tax != hasilTax {
		t.Errorf("SALAH! harusnya %.2f", hasilTax)
	}
}

func TestAmountCalculation(t *testing.T) {

	items, _ := TaxCalculate(1)
	t.Logf("Volume : %.2f", items.Amount)

	if items.Amount != hasilAmount {
		t.Errorf("SALAH! harusnya %.2f", hasilAmount)
	}
}
