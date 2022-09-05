package main
import "fmt"

func getFullName() (string, string) {
	return "Dicky", "Ardianto"
}

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
}

// _ underscore (jika tidak peduli dengan salah satu variabel)