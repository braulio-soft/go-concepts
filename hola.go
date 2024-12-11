package main

import (
	"fmt"
	"math"
	"rsc.io/quote"
	"strconv"
)

/*
// Define constant
const pi float32 = 3.1415926

// Define constant no necessary to use type
const pi2 = 3.14

// Another way to define constants
const (

	x = 100
	y = 0b1010 //Binary
	z = 0o12   //Octal
	w = 0xFF   //Hexadecimal

)
const (

	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado

)
*/
func main() {
	fmt.Println("Hola")
	fmt.Println(quote.Go())
	// Only accept positive value
	var integer uint = 0
	fmt.Println(integer)
	fmt.Println(math.MinInt64, math.MaxInt64)
	var valueBool bool = false
	fmt.Println(valueBool)
	fullname := "AlexRoel \t(alias \"roelcode\")\n"
	fmt.Println(fullname)
	//Another data type byte
	var a byte = 'a'
	fmt.Println(a)

	s := "hola"
	fmt.Println(s[0])
	//Another data type rune

	var r rune = '❤'
	fmt.Println(r)

	// ---------------------Const Default Values -------------------------
	var (
		defaultInt   int
		defaulUnit   uint
		defaulFloat  float32
		defaulBool   bool
		defaulString string
	)
	fmt.Println(defaultInt, defaulUnit, defaulFloat, defaulBool, defaulString)

	// -----------------------------------------------

	// ----------------Data types ------------------------------
	var integer16 int16 = 50
	var integer32 int32 = 100
	fmt.Println(integer16 + int16(integer32))

	string1 := "hola"
	i, result := strconv.Atoi(string1)
	fmt.Println(i+i, result)

	value1 := 42
	val := strconv.Itoa(value1)
	fmt.Println(val + val)
	// -----------------------------------------------

	// ---------------------Function FMT -------------------------

	name1 := "Braulio"
	age := 24
	fmt.Print("Another message")
	fmt.Printf("Hola , me llamo %s y tengo %d años. \n", name1, age)

	greeting := fmt.Sprintf("Hola , me llamo %s y tengo %d años.", name1, age)
	fmt.Println(greeting)

	var name10, name11 string
	var age10 int
	fmt.Print("Write your name : ")
	fmt.Scanln(&name10, &name11)
	fmt.Print("Write your age : ")
	fmt.Scanln(&age10)

	greeting2 := fmt.Sprintf("Hola , me llamo %s %s y tengo %d años.", name10, name11, age10)
	fmt.Println(greeting2)

	// -----------------------------------------------------------

	// --------------------- Math Operators ----------------------

	a1 := 10
	b := 4
	a1++
	a1--
	b++
	b--
	fmt.Println(a1 + b)
	fmt.Println(a1 - b)
	fmt.Println(a1 * b)
	fmt.Println(a1 / b)
	fmt.Println(a1 % b)

	// -----------------------------------------------------------

	/*
	   	//Define variable
	   	var firstName, lastName string
	   	var age int
	       // Other way
	   	var (
	   		firstName2 string
	   		lastName2 string
	   	    age2 int
	   	)
	*/
	//Another data
	var (
		firstName3 = "Alex"
		lastName3  = "Anderson"
		age3       = 12
	)
	//Last way these in and out functions
	var firstName4, lastName4, age4 = "Alex", "Anderson", 12
	// Define variable without using var only inside functions
	firstName5, lastName5, age5 := "Pedro", "Jhonson", 18

	fmt.Println(firstName3, lastName3, age3)
	fmt.Println(firstName4, lastName4, age4)
	fmt.Println(firstName5, lastName5, age5)
	/*
		fmt.Println(pi)
		fmt.Println(pi2)
		fmt.Println(x, y, z, w)
	*/

	/* Ejercicio: Calcule e imprima el área y el perímetro del triángulo
	Crear un programa que solicite al usuario que ingrese los lados de un
	triángulo rectángulo y luego calcule e imprima el área y el perímetro del triángulo.

	El programa debe:

	Solicitar al usuario que ingrese la longitud de los dos lados del triángulo rectángulo.

	Calcular la hipotenusa del triángulo usando el teorema de Pitágoras.

	Calcular el área del triángulo usando la fórmula base x altura / 2.

	Calcular el perímetro del triángulo sumando los lados.

	Imprimir el área y el perímetro del triángulo con dos decimales de precisión.

	El programa debe usar variables para almacenar los lados del triángulo,
	la hipotenusa, el área y el perímetro. También debe usar constantes para
	representar el número de decimales de precisión que se desean en la salida.

	Además, se deben utilizar funciones del paquete Math de Go para calcular
	la raíz cuadrada y cualquier otro cálculo matemático necesario.

	*/

	// -------------------------Solution Task 1-------------------------------------------------
	var (
		redneck1   int
		redneck2   int
		hypotenuse int
	)
	fmt.Print("Write rednecks for triangle : ")
	fmt.Scanln(&redneck1, &redneck2)
	hypotenuse = int(math.Sqrt(float64((redneck1 * redneck1) + (redneck2 * redneck2))))
	fmt.Println(redneck1, redneck2, hypotenuse)
	pe := redneck1 + redneck2 + hypotenuse
	ar := (redneck1 * redneck2) / 2
	perimeter := fmt.Sprintf("El perimetro del triangulo rectangulo de lados %d y %d es : %d", redneck1, redneck2, pe)
	area := fmt.Sprintf("El area del triangulo rectangulo de lados %d y %d es : %d", redneck1, redneck2, ar)
	fmt.Println(perimeter)
	fmt.Println(area)
	// --------------------------------------------------------------------------
}
