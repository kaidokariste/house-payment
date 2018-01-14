package main

import(
	"house-payment/common"
	"house-payment/routers"
	"github.com/codegangsta/negroni"
	"net/http"
	"log"
)

func main()  {
	//First thing would be init() in common package
	// Get the mux router object
	router := routers.InitRoutes()
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	//server := &http.Server{
	//	Addr:    ":9000",//common.AppConfig.Server,
	//	Handler: n,
	//}

	log.Println("Listening...", common.AppConfig.Server)
	//server.ListenAndServe()
	err := http.ListenAndServe(common.AppConfig.Server, n)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}