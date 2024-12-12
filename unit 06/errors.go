package main

import (
	"errors"
	"fmt"
	"strconv"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func main() {
	str := "123f"
	num, err := strconv.Atoi(str)
	fmt.Println(num)
	result, err := divide(10, 15)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println("Result : ", result)
	}
}
