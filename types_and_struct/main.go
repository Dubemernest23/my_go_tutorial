// package main

// import (
// 	"fmt"
// )

// type Vertex struct {
// 	name     string
// 	movement int
// 	isActive bool
// 	rating   float64
// }

// func main() {
// 	v := Vertex{"ali", 2, false, 2.4}
// 	v.isActive = true
// 	v.movement = 34
// 	v.name = "ahmed"
// 	v.rating = 4.0

// 	fmt.Println(v.isActive)
// 	fmt.Printf("The user's name is %v \n", v.name)
// }

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
