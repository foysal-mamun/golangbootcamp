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

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
func getLocation(city string) (string, string) {
	var region, continent string

	switch city {
	case "Dhaka", "dhaka":
		region, continent = "Bangladesh", "Asia"
	default:
		region, continent = "Unknown", "Unknown"
	}

	return region, continent
}
func getLongLat(city string) (long, lat int) {
	switch city {
	case "Dhaka", "dhaka":
		long, lat = 1, 2
	}

	return
}

type Artist struct {
	Name, Genre string
	Songs       int
}

func (a Artist) newRelease() int {
	a.Songs++
	return a.Songs
}
func (a *Artist) getRelease() int {
	a.Songs++
	return a.Songs
}
func main() {
	variable_typing()
	constants()
	printing()

	fmt.Println(add(10, 5))
	fmt.Println(sub(10, 5))

	region, continent := getLocation("Dhaka")
	fmt.Println(region, continent)

	long, lat := getLongLat("Dhaka")
	fmt.Println(long, lat)

	me := Artist{Name: "foysal", Genre: "Electro", Songs: 20}
	fmt.Println(me.Name, me.Genre, me.Songs)
	me.newRelease()
	fmt.Println(me.Name, me.Genre, me.Songs)
	me.getRelease()
	fmt.Println(me.Name, me.Genre, me.Songs)
}
