package main
import "fmt"

func random() interface{} {
	return true
}

func main() {
	var result interface{} = random()

	switch value  := result.(type){
	case string :
		fmt.Println("Value", value, "Is String")
	case int: 
		fmt.Println("Value", value, "Is int")
	default:
		fmt.Println("Unknown Type")
	}
}