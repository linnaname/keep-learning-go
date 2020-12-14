package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

type Animal struct {
	Name  string
	Order string
}

type Message struct {
	Name, Text string
}

func main() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

	var jsonBlob = []byte(`[
        {"Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)
	var animals []Animal
	e := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", e)
	}
	fmt.Println(animals)

	const jsonStream = `
        {"Name": "Ed", "Text": "Knock knock."}
        {"Name": "Sam", "Text": "Who's there?"}
        {"Name": "Ed", "Text": "Go fmt."}
        {"Name": "Sam", "Text": "Go fmt who?"}
        {"Name": "Ed", "Text": "Go fmt yourself!"}
    `
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
