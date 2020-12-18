package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Kind of PhoneNumber
type Kind byte

const (
	// MOBILE PhoneNumber
	MOBILE Kind = iota
	// HOME PhoneNumber
	HOME
	// WORK PhoneNumber
	WORK
)

// PhoneNumber of a person
type PhoneNumber struct {
	Number string
	Type   Kind
}

// Person info
type Person struct {
	ID     int
	Name   string
	Email  string
	Phones []PhoneNumber
}

// AddressBook of people
type AddressBook struct {
	People []Person
}

func main() {
	p := Person{
		ID:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []PhoneNumber{
			{Number: "555-4321", Type: HOME},
		},
	}
	bk := &AddressBook{}
	bk.People = append(bk.People, p)

	in, err := json.Marshal(bk)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(in)
	fmt.Println(string(in))
	fmt.Println("Length is", len(in))

	var book AddressBook
	if err := json.Unmarshal(in, &book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Println(book.People)
}
