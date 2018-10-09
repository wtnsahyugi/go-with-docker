package main

import (
	"fmt"
	"log"
	"net/http"

	"./jwt"

	migration "./migration"
	test "./pack/test"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("server")  // name of config file (without extension)
	viper.AddConfigPath("config/") // path to look for the config file in
	err := viper.ReadInConfig()    // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	port := viper.Get("port") //os.Getenv("PORT")

	setup := viper.Get("setup")

	if setup.(string) == "true" {
		migration.ProductMigration()
		migration.TaxMigration()
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware := jwt.Auth()

	// routing login
	r.POST("/login", authMiddleware.LoginHandler)

	router := r.Group("/V1")
	router.Use(authMiddleware.MiddlewareFunc())
	{
		router.GET("/product", test.ProductHandler)
		router.POST("/product/post", test.CreateProductHandler)
		router.GET("/refresh_token", authMiddleware.RefreshHandler)
		router.GET("/taxcalculation/:id", test.TaxCalculationHandler)
		// router.POST("/hello/customer/post", hello.CreateCustomerHandler)
		// router.POST("/hello/customer/upload", hello.UploadCustomerHandler)
	}

	fmt.Printf(port.(string))
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
	//http.ListenAndServe(":"+port.(string), r)

}
