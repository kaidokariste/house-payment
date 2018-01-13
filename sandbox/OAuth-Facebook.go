package main

import (
	"io/ioutil"
	"os"
	"html/template"
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth"
	//"github.com/gorilla/pat"
	"github.com/markbates/goth/providers/facebook"
	"github.com/gorilla/mux"
)

// Struct for parsing JSON configuration.
type Configuration struct {
	FacebookKey    string `json:"fbkey"`
	FacebookSecret string `json:"fbsecret"`
	FacebookCallback string `json:"fbcallback"`
}

var conf Configuration

func init() {
	//Use Readfile to read cred.json into Configuration struct
	file, err := ioutil.ReadFile("sandbox/creds.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &conf)
}

func callbackAuthHandler(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	fmt.Println(user.UserID, user.Name, user.AccessToken, user.ExpiresAt)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	t, _ := template.New("userinfo").Parse(userTemplate)
	t.Execute(res, user)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	t, _ := template.New("index").Parse(indexTemplate)
	t.Execute(res, nil)
}

func main() {
	//Register OAuth2 providers with Goth
	goth.UseProviders(
		facebook.New(conf.FacebookKey, conf.FacebookSecret, conf.FacebookCallback),
	)
	//Routing using Pat package
	//r := pat.New()
	//r.Get("/auth/{provider}/callback", callbackAuthHandler)
	//r.Get("/auth/{provider}", gothic.BeginAuthHandler)
	//r.Get("/", indexHandler)

	//Routing using mux package
	r := mux.NewRouter()
	r.HandleFunc("/auth/{provider}/callback",callbackAuthHandler).Methods("GET")
	r.HandleFunc("/auth/{provider}", gothic.BeginAuthHandler).Methods("GET")
	r.HandleFunc("/", indexHandler).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}

	//View templates
	var indexTemplate = `
<p><a href="/auth/facebook">Log in with Facebook</a></p>
`

	var userTemplate = `
<p>Name: {{.Name}}</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
`