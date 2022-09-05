// recover adalah function yang digunakan untuk menangkap data panic
// dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan

package main
import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
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