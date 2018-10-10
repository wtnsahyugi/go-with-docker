- Berikut adalah langkah-langkah untuk mengambil data perhitungan bill user menggunakan golang

1. request token untuk dijadikan bearer authorization dengan hit api http://localhost/login dengan formdata 
{
    "username":"admin",
    "password":"admin"
}

2. setelah token didapat maka untuk mengambil data bill dengan hit api http://localhost/V1/taxcalculation/{id} , id diisi dengan integer id dari product id

- Desain database 
#products
*id
-name
-price
-tax_id

#taxes
*id
-name
-refundable

- Untuk mengeksekusi unit test bisa dicoba dengan langkah-langkah berikut:
1. masuk ke go bash dengan mengeksekusi docker-compose exec go bash
2. lalu masuk ke folder /go/src/test2/pack/test/repo
3. lalu eksekusi go test taxcalculation_repo.go taxcalculation_repo_test.go -v