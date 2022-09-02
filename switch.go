package main
import "fmt"

func main()  {
	name := "Dicky"

	switch name {
		case "Dicky":
			fmt.Println("Hallo", name)
		case "Ardianto":
			fmt.Println("Hallo", name)
		default:
			fmt.Println("Hallo Kenalan dong")
	}
}