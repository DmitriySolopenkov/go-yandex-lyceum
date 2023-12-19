package main

import (
	"os"
	"text/template"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func main() {
	t := template.Must(template.ParseFiles("template.go.tmpl"))
	data := struct {
		Package string
		Struct  string
		User    string
	}{
		Package: "db",
		Struct:  "UserRepository",
		User:    "User",
	}
	t.Execute(os.Stdout, data)
}
