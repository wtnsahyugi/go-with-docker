package repo

import (
	"../../../config"
	"../model"
)

// type listData []map[string]string

// func initData(param string) string {
// 	var listfunct listData

// 	listfunct = []map[string]string{
// 		{"value": "1", "function": "foodTaxCalculate"},
// 		{"value": "2", "function": "tobbacoTaxCalculate"},
// 		{"value": "3", "function": "entertainTaxCalculate"},
// 	}

// 	return listfunct.find(param)

// }

// func (data listData) find(byValue string) string {
// 	var found string

// 	found = ""
// 	for _, val := range data {
// 		if byValue == val["value"] {
// 			found = val["function"]
// 			break
// 		}
// 	}
// 	return found
// }

var funcMap = map[string]interface{}{
	"1": foodTaxCalculate,
	"2": tobbacoTaxCalculate,
	"3": entertainTaxCalculate,
}

func callDynamically(name string, args ...interface{}) int {
	return funcMap[name].(func(model.GetProduct) int)(args[0].(model.GetProduct))

}

func foodTaxCalculate(mod model.GetProduct) int {

	taxtotal := (mod.Price * 10) / 100

	return taxtotal
}

func tobbacoTaxCalculate(mod model.GetProduct) int {

	taxtotal := 10 + ((mod.Price * 2) / 100)

	return taxtotal
}

func entertainTaxCalculate(mod model.GetProduct) int {

	if mod.Price < 100 {
		return 0
	}

	taxtotal := (mod.Price - 100) / 100

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

	return &mod, nil
}
