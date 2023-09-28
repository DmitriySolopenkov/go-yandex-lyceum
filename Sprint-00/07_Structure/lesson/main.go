package main

import (
	"fmt"

	"github.com/DmitriySolopenkov/go-yandex-lyceum/07_Structure/lesson/geometry"
	"github.com/DmitriySolopenkov/go-yandex-lyceum/07_Structure/lesson/person"
	"github.com/DmitriySolopenkov/go-yandex-lyceum/07_Structure/lesson/students"
)

// Приведение типов
func PrintType(x interface{}) {
	fmt.Printf("%v, %T\n", x, x)
	// switch x.(type) {
	// case string:
	// 	fmt.Println("string")
	// case int:
	// 	fmt.Println("int")
	// case bool:
	// 	fmt.Println("bool")
	// default:
	// 	fmt.Println("unknown")
	// }
}

func main() {
	student1 := students.Student{}
	fmt.Println(student1)
	rect := geometry.Rectangle{Width: 10, Height: 5}
	fmt.Println(rect.Area())
	scaledW, scaledH := rect.Scale(2.0)
	fmt.Println(scaledW)
	fmt.Println(scaledH)
	var h person.Human = person.Person{Name: "John", Age: 25}
	h.SayHello()
	fmt.Println("S =", geometry.CalculateArea(rect))

	// Приведение типов
	PrintType("hello") // string
	PrintType(42)      // int
	PrintType(true)    // bool
	PrintType(3.14)    // unknown

}

//width: 10, height: 5
