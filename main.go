package main

import(
	"house-payment/common"
	"fmt"

)

func main()  {

	//Load in the configuration data
	common.StartUp()
	// Get the mux router object
	//router := routers.InitRoutes()

	//Test Parsing configuration
	fmt.Println(common.AppConfig.Database)
}