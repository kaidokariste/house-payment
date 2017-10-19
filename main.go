package main

import(
	"house-payment/common"
	"fmt"
)

func main()  {

	//Load in the configuration data
	common.StartUp()

	//Test Parsing configuration
	fmt.Println(common.AppConfig.Database)
}