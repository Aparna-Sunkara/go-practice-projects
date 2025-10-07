package main

import (
	"fmt"
	"strings"
	
)

type User struct {
	UserID        string
	PhoneNumber   string
	Contacts      map[string]*Contact 
	Messages      []Message
}


type Contact struct {
	Name        string
	PhoneNumber string
}

type Message struct {
	Sender    string
	Receiver  string
	
}
var allUsers = make(map[string]*User)
var currentUser *User
func main() {
	
}