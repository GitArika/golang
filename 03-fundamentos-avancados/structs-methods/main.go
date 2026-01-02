package main

import (
	"encoding/json"
	"fmt"
	"structs-methods/bar"
)

type MinhaString string

type User struct{
	Name string `json:"name"`
	ID uint64 `json:"id"`
	bar.Bar // embed
}

func (u User) Foo() {
	fmt.Println(u.Name)
} 

func (u *User) UpdateName(newName string) {
	u.Name = newName
}

func main() {
	user := User{Name: "Ariel", ID: 10}
	fmt.Println(user)
	user.Foo()
	user.UpdateName("GitArika")
	user.Foo()
	user.DoThing()

	res, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

