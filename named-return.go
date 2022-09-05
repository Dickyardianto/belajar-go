package main
import "fmt"

func getFullName2() (firstName string, middleName string, lastName string) {
	firstName = "Dicky"
	middleName = "Ardianto"
	lastName = "Saja"

	return
}

func main() {
	firstName, middleName, lastName := getFullName2() 
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}