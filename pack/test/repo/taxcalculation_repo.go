package repo

import (
	"../../../config"
	"../model"
)

var funcMap = map[string]interface{}{
	"1": foodTaxCalculate,
	"2": tobbacoTaxCalculate,
	"3": entertainTaxCalculate,
}

func callDynamically(name string, args ...interface{}) float32 {
	return funcMap[name].(func(model.GetProduct) float32)(args[0].(model.GetProduct))

}

func foodTaxCalculate(mod model.GetProduct) float32 {

	taxtotal := (mod.Price * 10) / 100

	return taxtotal
}

func tobbacoTaxCalculate(mod model.GetProduct) float32 {

	taxtotal := 10 + ((mod.Price * 2) / 100)

	return taxtotal
}

func entertainTaxCalculate(mod model.GetProduct) float32 {

	if mod.Price < 100 {
		return 0
	}

	taxtotal := (mod.Price - 100) * 1 / 100

	return taxtotal
}

func TaxCalculate(prodId int) (*model.GetProduct, error) {
	var mod model.GetProduct

	db, err := config.GetConn()

	if err != nil {
		return nil, err
	}

	db.Table("products").Select("products.name, products.price, taxes.tax_name, taxes.refundable, products.tax_id").Joins("left join taxes on products.tax_id = taxes.id").Where("products.id = ?", prodId).Scan(&mod)

	reserveTax := callDynamically(mod.TaxID, mod)

	mod.Tax = reserveTax
	mod.Amount = mod.Price + mod.Tax

	return &mod, nil
}
