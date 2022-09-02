package main
import "fmt"

func main() {
	name := "Ardianto"

	if name == "Dicky" {
		fmt.Println("Hallo",name)
	}else if name == "Ardianto" {
		fmt.Println("Hallo", name)
	}else {
		fmt.Println("Tidak ada data")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	}else {
		fmt.Println("Nama Sudah Benar")	
	}
}