//interface biasanya digunakan sebagai kontrak
// adalah tipe data abstrak

package main
import "fmt"

type HasName interface {
	GetName() string
}

func sayHayy(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var dicky Person
	dicky.Name = "Dicky"
	sayHayy(dicky)

	cat := Animal{
		Name: "Kucing",
	}
	sayHayy(cat)
}