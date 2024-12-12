package main

import "fmt"

func main() {
	// [...] when you don't know how many space have the matrix
	var matrix = [...]int{10, 20, 30, 40, 50}
	//function len() to get size for the matrix
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	// Another way to iterate values in a matrix . you can use _ , value or index , _
	for index, value := range matrix {
		fmt.Printf("Indice = %d, value = %d\n", index, value)
	}

	//BiDimensional matrix
	var matrix2 = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix2)
	for index, value := range matrix2 {
		fmt.Printf("Index = %d, value = %d\n", index, value)
	}

	//------------------------ Part for Array (Slices) -----------------------------

	daysOfWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	weekDays := daysOfWeek[0:5]
	fmt.Println(daysOfWeek)
	fmt.Println(weekDays)
	weekDays = append(weekDays, "Saturday")
	fmt.Println(weekDays)
	weekDays = append(weekDays[:1], weekDays[2:]...)
	fmt.Println(weekDays)
	fmt.Println(len(weekDays))
	fmt.Println(cap(weekDays))

	// How to create slices using method make
	names := make([]string, 5, 10)
	names[0] = "Alex"
	fmt.Println(names)

	//Other case
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5)
	copy(slice2, slice1)

	fmt.Println(slice1)
	fmt.Println(slice2)

	// ------------------------- Maps ----------------------------------
	colors := map[string]string{"red": "#FFOO00", "blue": "#0000FF", "green": "#00FF00"}
	fmt.Println(colors)
	colors["black"] = "#000000"
	fmt.Println(colors)
	valor, ok := colors["white"]
	if ok {
		fmt.Println(valor)
	} else {
		fmt.Println("These value do not exist")
	}
	delete(colors, "black")
	fmt.Println(valor, ok)

	for key, value := range colors {
		fmt.Println(key, value)
	}
}
