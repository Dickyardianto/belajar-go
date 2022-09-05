package main
import "fmt"

type BlackList func(string) bool 

func registerUser(name string, blackList BlackList){
	if blackList(name) {
		fmt.Println("Kamu di block", name)
	}else {	
		fmt.Println("Welcome", name)
	}
}

func main() {
	blackList := func(name string) bool {
		return name == "BUCIN TAI"
	}

	registerUser("admin", blackList)
	registerUser("BUCIN TAI", blackList)

	registerUser("root", func(name string) bool {
		return name == "BUCIN TAI"
	})

	registerUser("root", func(name string) bool {
		return name == "root"
	})
}