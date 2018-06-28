package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type User struct {
	email string
	name  string
}

type Database struct {
	users []User
}

//var Data Database = Database{users: []User{}}
var Data Database

func find(pemail string, ch chan *User, sidx, eidx int, name string) {
	for i := sidx; i < eidx; i++ {
		user := &Data.users[i]
		if user.email == pemail {
			//found = true
			fmt.Println("\n\n\nFound By worker:", name)
			ch <- user
		}
	}

}

func main() {
	pemail := os.Args[1]
	var user User
	var suser *User

	for i := 0; i < 500; i++ {
		user.name = "User" + strconv.Itoa(i)
		user.email = "User" + strconv.Itoa(i) + "@example.com"
		Data.users = append(Data.users, user)
	}
	fmt.Println(Data.users)
	user = User{}
	ch := make(chan *User)
	//chout := make(chan *User)
	go find(pemail, ch, 0, 100, "#1")
	go find(pemail, ch, 100, 200, "#2")
	go find(pemail, ch, 200, 300, "#3")
	go find(pemail, ch, 300, 400, "#4")
	go find(pemail, ch, 400, 500, "#5")
	select {
	case suser = <-ch:
		fmt.Println("User found. User Name:", suser.name, "  email:", suser.email, "\n")
	case <-time.After(1 * time.Second):
		fmt.Println("User not found.")
	}

}
