package main
import "fmt"

func main() {
	counter := 0

	//sebuah blok scope itu bisa berinteraksi dengan yang lain
	increment := func() {
		fmt.Println("Increment")
		counter++
	}
	
	increment()
	increment()
	fmt.Println(counter)
}