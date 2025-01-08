package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		div3 := false
		div5 := false
		div3Msg := ""
		div5Msg := ""
		if i%3 == 0 {
			div3 = true
			div3Msg = "Pin "
		}
		if i%5 == 0 {
			div5 = true
			div5Msg = "Pan"
		}

		if div3 || div5 {
			fmt.Printf("%s%s\n", div3Msg, div5Msg)
		}
	}
}
