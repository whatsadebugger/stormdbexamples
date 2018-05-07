package main

import (
	"fmt"
	"github.com/asdine/storm"
	"time"
)

// User holds basic account information
type User struct {
	ID        int    `storm:"id"`     // primary key
	Group     string `storm:"index"`  // this field will be indexed
	Email     string `storm:"unique"` // this field will be indexed with a unique constraint
	Name      string // this field will not be indexed
	Age       int    `storm:"index"`
	CreatedAt time.Time
}

func main() {
	fmt.Print("Welcome to stormdb examples\n")
	db, err := storm.Open("examples.db")
	PanicIfError(err)

	user := User{
		ID:        13,
		Group:     "student",
		Email:     "ahmad@goodschool.net",
		Name:      "Ahmad",
		Age:       24,
		CreatedAt: time.Now(),
	}

	user1 := User{
		ID:        14,
		Group:     "student",
		Email:     "joe@goodschool.net",
		Name:      "Joseph",
		Age:       21,
		CreatedAt: time.Now(),
	}

	user2 := User{
		ID:        15,
		Group:     "teacher",
		Email:     "john@goodschool.net",
		Name:      "John Titor",
		Age:       18,
		CreatedAt: time.Now(),
	}

	PanicIfError(db.Save(&user))
	db.Save(&user1)
	db.Save(&user2)

	// example of finding a user by email
	var lookup []User
	db.Find("Email", "john@goodschool.net", &lookup)
	fmt.Printf("%+v\n", lookup)

	// all users in reverse by ID
	var all []User
	db.All(&all, storm.Reverse())
	for _, v := range all {
		fmt.Printf("%+v\n", v)
	}
}

// PanicIfError will panic if err != nil
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
