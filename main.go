package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

const (
	port = ":8080"
	host = "http://localhost"
)

func main() {

	// Variant 1

	tmpl1, err := template.New("test1").Parse("the string is {{.}}\n") // {{.}} is the default variable

	if err != nil {
		fmt.Println(err)
	}
	if tmpl1.Execute(os.Stdout, 123) != nil { // 123 is the default variable
		fmt.Println(err)
	}

	// Variant 2

	tmpl2, err := template.New("test2").Parse("the string is {{.Color}}\n")
	if err != nil {
		fmt.Println(err)
	}
	a := struct {
		Color string
		Age   int
	}{"blue", 11}
	if tmpl2.Execute(os.Stdout, a) != nil {
		fmt.Println(err)
	}

	// Variant 3
	type stock struct {
		Material string
		Count    uint
	}
	sweaters := stock{"wool", 17}
	tmp3, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmp3.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

	// variant 4

	Var4()
}

// varian 4  Output information not to the terminal but to the page

// https://metanit.com/go/web/2.2.php
func Var4() {
	type ViewData struct {
		Title string
		Users []string
	}

	data := ViewData{
		Title: "Users List",
		Users: []string{"Tom", "Bob", "Sam"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // handles requests to the root path of the server and executes the function inside the brackets for each request
		tmp4, _ := template.ParseFiles("templates/index.html")
		tmp4.Execute(w, data) // data is the default variable for the template file index.html
	})

	fmt.Println("Server is listening...")                                                 // prints to console
	log.Printf("Starting application on port%s, press CTRL + c to shut it down.\n", port) // prints to console
	log.Println("Open:", "\033[1;32m"+host+port+"\033[0m")

	if err := http.ListenAndServe(port, nil); err != nil { //starts server on port 8080 and listens for requests
		log.Fatal(err)
	}
	http.ListenAndServe(":8080", nil)
}
