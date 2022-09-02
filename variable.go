/* variabel di golang
satu variable itu unik di golang
ada variable yang tidak dipake, dia akan protes
var tidak wajib asal kita langsung menginisialisasi datanya
*/

package main
import "fmt"

func main() {
	var nama string
	
	nama = "Dicky Ardianto"
	fmt.Println(nama)

	// nama = true
	// fmt.Println(nama)

	country := "Jakarta"
	fmt.Println(country)

	country = "Bekasi"
	fmt.Println(country) 
	
	// variabel multiple
	
	var (
		firstname = "Dicky"
		lastname = "Ardianto"
	)
	fmt.Println(firstname) 
	fmt.Println(lastname) 

}




