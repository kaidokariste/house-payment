package main

import(
	"encoding/json"
	"strings"
	"io"
	"log"
	"fmt"
)

func ExampleDecoder() {

	const jsonStream = `

		{"townName": "Tartu", "currentPopulation": 5}
		{"townName": "Tallinn", "currentPopulation": 8}
		{"townName": "Riia", "currentPopulation": 88}
		{"townName": "Pärnu", "currentPopulation": 9}
		{"townName": "Narva", "currentPopulation": 11}
	`

	type Message struct {
		TownName string
		CurrentPopulation int
	}

	dec := json.NewDecoder(strings.NewReader(jsonStream))

	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break

		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.TownName, m.CurrentPopulation)
	}
}

func main(){
	ExampleDecoder()
}