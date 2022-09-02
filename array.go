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

	fmt.Println(len(names))
	fmt.Println(len(values))

	// slice 
	var months = [...]string {
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	
	months[5] = "Diubah"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)
	
	var slice3 = append(slice2, "Dicky")
	fmt.Println(slice3)
	
}