package main

import (
	"testing"
	"fmt"
)

func TestCalc(t *testing.T) {
	fmt.Printf("d = 2     : %d\n", Calc("d = 2"))
	fmt.Printf("3         : %d\n", Calc("3"))
	fmt.Printf("3 * 4     : %d\n", Calc("3 * 4"))
	fmt.Printf("b = 5 * 6 : %d\n", Calc("b = 5 * 6"))
	fmt.Printf("2 / 2     : %d\n", Calc("2 / 2"))
	fmt.Printf("c = 4 / 2 : %d\n", Calc("c = 4 / 2"))
	fmt.Printf("4 / 2 + 3 : %d\n", Calc("4 / 2 + 3"))
	fmt.Printf("3 + 4 / 2 : %d\n", Calc("3 + 4 / 2"))
	fmt.Printf("4 / 2 * 2 : %d\n", Calc("4 / 2 * 2"))
}
