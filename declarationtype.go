package main
import "fmt"

func main() {
	type NoKTP string

	var noKTPDicky NoKTP = "1029109201920192012"
	fmt.Println(noKTPDicky)
}