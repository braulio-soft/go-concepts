package main

import (
	"fmt"
	"os"
)

func main() {
	//defer to pas the process to the END if we pass all . it's going to be 213
	defer fmt.Println(3)
	defer fmt.Println(1)
	defer fmt.Println(2)
	//Example using DEFER
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err3 := file.Write([]byte("Hello World\n"))
	if err3 != nil {
		fmt.Println(err3)
		return
	}

}
