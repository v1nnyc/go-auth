package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	First string
}

func main() {
	p1 := Person{
		First: "vincent",
	}

	p2 := Person{
		First: "gabe",
	}

	xp := []Person{p1, p2}

	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("json printing: %v\n", string(bs))

	xp2 := []Person{}

	err = json.Unmarshal(bs, &xp2)

	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("xp2: %v\n", xp2)

}
