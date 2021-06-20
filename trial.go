package gomc

import (
	"fmt"
)

func TrialMain(args []string) {
	fmt.Println(args)
	var a int = 100
	var b float64
	var c string
	var d bool
	fmt.Printf("a: %d, b:%f, c:%s, d:%t\n", a, b, c, d)
}