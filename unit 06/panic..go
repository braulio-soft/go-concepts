package main

import "fmt"

func division(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	//validateZero(b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)

}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("divisor is zero")
	}
}

func main() {
	division(12, 1)
	division(12, 2)
	division(12, 0)
	division(120, 15)

}
