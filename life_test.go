// Copyright 2017 Aleksandr Demeshko. All rights reserved.

// conway project
// life_test.go
package conway

import (
	"fmt"
	//	"testing"
)

func Example_singleCell() {
	var p = Population{cells: map[Cell]int{
		Cell{0, 0}: 0,
	},
		popNumber: 0,
	}
	p.Next()
	fmt.Println(p)
	// Output: {map[] 1}
}

func Example_twoCells() {
	var p = Population{cells: map[Cell]int{
		Cell{0, 0}: 0,
		Cell{0, 1}: 0,
	},
		popNumber: 0,
	}
	p.Next()
	fmt.Println(p)
	// Output: {map[] 1}
}

func Example_blinker() {
	var p = Population{cells: map[Cell]int{
		Cell{0, 0}:  0,
		Cell{0, 1}:  0,
		Cell{0, -1}: 0,
	},
		popNumber: 0,
	}
	p.SaveToFile("blinker0.log")
	p.Next()
	p.SaveToFile("blinker1.log")
	fmt.Println(len(p.cells))
	// Output: 3
}
