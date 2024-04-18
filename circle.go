package shapes

import "fmt"

type Circle struct {
	Center Point
	Radius float32
}

func (c Circle) Print() { fmt.Println(c) }
