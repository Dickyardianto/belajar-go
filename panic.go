package main
import "fmt"

// panic function adalah yang bisa digunakan untuk menghentikan program
// biasa dipanggil ketika terjadi error saat program berjalan

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
}