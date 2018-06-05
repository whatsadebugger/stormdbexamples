package main

import (
	"fmt"
	"time"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
)

// User holds basic account information
type User struct {
	ID        int
	Group     string `storm:"index"`  // this field will be indexed
	Email     string `storm:"unique"` // this field will be indexed with a unique constraint
	Name      string // this field will not be indexed
	Age       int    `storm:"index"`
	CreatedAt time.Time
}

// School is a great place
type School struct {
	ID      int
	Name    string `storm:"index"`
	City    string `storm:"index"`
	Founded int    `storm:"index"`
}

func main() {
	fmt.Print("Welcome to stormdb examples\n")
	db, err := storm.Open("examples.db")
	PanicIfError(err)

	user0 := User{
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

	user3 := User{
		ID:        16,
		Group:     "student",
		Email:     "derk@goodschool.net",
		Name:      "derk",
		Age:       22,
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

	school0 := School{
		ID:      1,
		Name:    "Jasper",
		City:    "Plano",
		Founded: 1996,
	}

	school1 := School{
		ID:      2,
		Name:    "Rice",
		City:    "Plano",
		Founded: 1999,
	}

	db.Save(&school0)
	db.Save(&school1)

	db.Save(&user0)
	db.Save(&user1)
	db.Save(&user2)
	db.Save(&user3)

	// example of finding a user by email
	var users []User
	db.Find("Email", "john@goodschool.net", &users)
	fmt.Printf("Found User with matching email:\n %+v\n", users)

	// Find a student by ID
	err = db.Find("ID", 13, &users)
	PanicIfError(err)
	fmt.Printf("Found Student with matching ID:\n %+v\n", users)

	// Find all schools using the name index and in reverse order
	var foundSchools []School
	err = db.AllByIndex("Name", &foundSchools, storm.Reverse())
	PanicIfError(err)

	for i, v := range foundSchools {
		fmt.Printf("School #%v:\n %+v\n", i, v)
	}

	// all users in reverse by ID
	fmt.Println("All users reversed:")
	var all []User
	db.All(&all, storm.Reverse())
	for _, v := range all {
		fmt.Printf(" %+v\n", v)
	}

	// find students using advanced queries
	db.Select(q.Eq("Group", "student")).Limit(5).Reverse().OrderBy("Age").Find(&users)
	fmt.Printf("Students Group: \n %+v\n", users)

	db.Select(q.Eq("Group", "student")).Limit(5).Reverse().OrderBy("Email").Find(&users)
	fmt.Printf("Students Group: \n %+v\n", users)

	db.Select(q.Eq("Group", "student")).Limit(5).Reverse().OrderBy("ID").Find(&users)
	fmt.Printf("Students Group: \n %+v\n", users)

}

// PanicIfError will panic if err != nil
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
