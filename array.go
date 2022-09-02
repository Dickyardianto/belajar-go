// array , kumpulan data dengan tipe yang sama

package main
import "fmt"
func main() {
	var names [3]string
	names[0] = "Dicky"
	names[1] = "Ardianto"
	names[2] = "S"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		91,
		92,
	}

	fmt.Println(names)
	fmt.Println(values)
}