package main

import (
	"fmt"
	"projects/datatypes"
	"projects/math"
)

func main() {
	fmt.Println("This is main package")

	math.CallMe()
	math.Add(2, 3)
	math.Multiply(3, 2)

	datatypes.DataTypes()
}
