package main

import (
	"fmt"
	"github.com/asdine/storm"
	"time"
)

// User holds basic account information
type User struct {
	ID        int       `storm:"id"`     // primary key
	Group     string    `storm:"index"`  // this field will be indexed
	Email     string    `storm:"unique"` // this field will be indexed with a unique constraint
	Name      string    // this field will not be indexed
	Age       int       `storm:"index"`
	CreatedAt time.Time `storm: "index"`
}

func main() {
	fmt.Print("Welcome to stormdb examples\n")

	user := User{
		ID:        13,
		Group:     "student",
		Email:     "ahmad@goodschool.net",
		Name:      "Ahmad",
		Age:       24,
		CreatedAt: time.Now(),
	}

	// PanicIfError(storm.Save(&user))
}

// PanicIfError will panic if err != nil
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
