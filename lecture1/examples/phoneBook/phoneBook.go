package main

import (
	"errors"
	"fmt"
)

// simply run `go run ./lecture1/examples/phoneBook/phoneBook.go` from the root of the repository

type phoneBook struct {
	namesToPhones map[string]string
}

func NewPhoneBook() *phoneBook {
	return &phoneBook{namesToPhones: map[string]string{}}
}

func (p *phoneBook) Add(name string, phone string) error {
	_, found := p.namesToPhones[name]
	if found {
		return errors.New("name already exists")
	}
	p.namesToPhones[name] = phone
	return nil
}

func (p *phoneBook) Remove(name string) {
	delete(p.namesToPhones, name)
}

func (p *phoneBook) List() {
	fmt.Println("Listing phone book")
	for name, phone := range p.namesToPhones {
		output := fmt.Sprintf("[name = %s, phone = %s]; ", name, phone)
		fmt.Print(output)
	}
	fmt.Println()
}

func main() {
	book := NewPhoneBook()
	book.List()

	if err := book.Add("alice", "+1235648"); err != nil {
		fmt.Println("failed to add Alice to phone book: ", err)
	}
	if err := book.Add("bob", "+9999922222"); err != nil {
		fmt.Println("failed to add Bob to phone book: ", err)
	}
	if err := book.Add("kate", "+777777777"); err != nil {
		fmt.Println("failed to add Kate to phone book: ", err)
	}

	if err := book.Add("alice", "+747474747"); err != nil {
		fmt.Println("failed to add alice: ", err)
	}

	book.List()
	book.Remove("bob")
	book.List()
}
