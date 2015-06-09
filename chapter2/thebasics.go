package main

import (
	"fmt"
)

func variable_typing() {

	var (
		name     string
		age      int
		location string
	)
	var (
		city, country string
		zip           int
	)

	var comment string
	var rating int

	var (
		name1 string = "Prince Obryn"
		age1  int    = 12
	)
	var (
		location1 = "Dhaka"
	)

	var (
		city1, zip1 = "Dhaka", 1234
	)

	city2, zip2 := "Berlin", 7890

	displayAll := func() {
		fmt.Println(name, age, location, city, country, zip, comment, rating, name1, age1, location1, city1, zip1, city2, zip2)
	}
	displayAll()
}

func constants() {
	const Pi = 3.14
	const (
		StatusOk      = 200
		StatusCreated = 201
	)

	fmt.Println(Pi, StatusOk, StatusCreated)
}

func printing() {
	cylonModel := 6
	fmt.Println(cylonModel)

	name := "Foysal"
	msg := fmt.Sprintf("Number %d", 6)
	fmt.Printf("%s is also known as %s\n", name, msg)
}

func main() {
	variable_typing()
	constants()
	printing()
}
