// namanya adalah for loops(cuma ada 1 perulangan)

package main
import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	// for loop ada 2 statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}
	
	slice := []string{"Dicky", "Ardianto", "Alias"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
