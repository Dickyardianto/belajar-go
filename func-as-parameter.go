package main 
import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "BUCIN TAI" {
		return "..."
	}else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Dicky", spamFilter)
	sayHelloWithFilter("BUCIN TAI", spamFilter)
}