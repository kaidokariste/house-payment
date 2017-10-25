package main

import(
	"house-payment/common"
	//"fmt"

	"house-payment/routers"
	//"github.com/codegangsta/negroni"
	"net/http"
	"log"
)

func main()  {

	//Load in the configuration data
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	log.Print("Mux router object", router)
	// Create a negroni instance
	//n := negroni.Classic()
	//n.UseHandler(router)

	//server := &http.Server{
	//	Addr:    ":9000",//common.AppConfig.Server,
	//	Handler: n,
	//}
	log.Println("Listening...")
	//server.ListenAndServe()
	err := http.ListenAndServe(":9000", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


	//Test Parsing configuration
	//fmt.Println(common.AppConfig.Database)
}