package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("It's morning ")
	} else if t.Hour() > 12 && t.Hour() < 17 {
		fmt.Println("It's afternoon")
	} else {
		fmt.Println("It's even")
	}

	//------------------------------- Switch Case -----------------------------

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run Windows")
	case "linux":

		fmt.Println("Go run Linux")
	case "darwin":
		fmt.Println("Go run macOS")
	default:
		fmt.Println("Go run Other OS")
	}

	//------------------------------- For Case -----------------------------

	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		if i == 2 {
			continue
		}
		time.Sleep(1 * time.Second)
	}
	//------------------------------- Functions -----------------------------
	greeting := hello("Braulio")
	fmt.Println(greeting)
	sum, mul := cal(12, 15)
	fmt.Println(sum, mul)

}

func hello(name string) string {
	return "Hello " + name
}

func cal(val1, val2 int) (int, int) {
	sum := val1 + val2
	mul := val1 * val2
	return sum, mul
}
