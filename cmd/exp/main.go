package main

import (
	"html/template"
	"os"
)

type UserMeta struct {
	Visits int
}

type User struct {
	Name string
	Meta UserMeta
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "moris borris",
		Meta: UserMeta{
			Visits: 5,
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
