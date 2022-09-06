// adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnnya dalam satu kesatuan
// struct adalah kumpulan dari field
package main 
import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func main() {

	var customer Customer
	customer.Name = "Dicky"
	customer.Address = "Indonesia"
	customer.Age = 30

	// struct literals

	customer2 := Customer {
		Name: "Dicky Ardianto",
		Address: "Indonesia",
		Age: 23,
	}

	fmt.Println(customer)
	fmt.Println(customer2)
}