package main

import "fmt"

func main() {

	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	var firstName string = "John"

	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	lastName := "wick"

	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")
}
