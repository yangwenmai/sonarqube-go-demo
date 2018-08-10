package main

import (
	"fmt"

	"gitlab.kenny.com/demo/math"
)

// Unused 无用的
var Unused int

func main() {
	m := math.New(10, 2)

	result := m.Divide()
	fmt.Printf("%d/%d=%d\n", m.GetA(), m.GetB(), result)

	mathMap, err := m.ToMap()
	if err != nil {
		panic(err)
	}
	fmt.Printf("mathMap: %+v\n", mathMap)

	fmt.Printf("Unused: %d\n", Unused)
}
