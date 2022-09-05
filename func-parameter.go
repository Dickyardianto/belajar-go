package main

import "fmt"

func sayHelloTo(firstName string, lastName string, age int) {
	fmt.Println("Hello", firstName, lastName, "Umur", age)
}

func main() {
	sayHelloTo("Dicky", "Ardianto", 23)
}