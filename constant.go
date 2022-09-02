// variabel yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai

package main
import "fmt"

func main() {
	const firstname string = "Dicky"
	const lastname = "Ardianto"
	const value = 1000

	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(value)

	// constant multiple/ sekali di deklarasikan datanya maka tidak akan bisa diubah lagi
	const (
		a string = "Dicky"
		b = "Ardianto"
		c = 1000
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}