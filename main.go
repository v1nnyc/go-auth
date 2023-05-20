package main

import (
	// "encoding/json"
	// "fmt"
	// "log"
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	First string
}

func main() {
	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := Person{
		First: "vincent",
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(p1)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
func bar(w http.ResponseWriter, r *http.Request) {
	var xp []Person
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&xp)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("xp: %v\n", xp)
}
