package main

import "fmt"

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println(`HOLA MUNDO`)
	fmt.Println(`PI`, pi)
	fmt.Println(`PI2`, pi2)

	//Declaracion de variables
	base := 12
	var altura int = 14
	var nombre string = "yasniel"
	var edad int = 27

	fmt.Println(base, altura, nombre, edad)

	// Zero values, lo que seria null en otros lenguajes
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("areaCuadrado: ", areaCuadrado)

	//---------------------------------------------OPERADORES---------------------------------------------
	x := 10
	y := 50

	// Suma
	suma := x + y
	fmt.Println("suma: ", suma)

	// Resta
	resta := y - x
	fmt.Println("resta: ", resta)

	// Multiplicacion
	multiplicacion := x * y
	fmt.Println("multiplicacion: ", multiplicacion)

	//Division
	division := x / y
	fmt.Println("division: ", division)

	// Calculo del resto
	resto := y % x
	fmt.Println("resto: ", resto)

	//Incremeto +1
	x++
	fmt.Println("incremento mas 1 de 10: ", x)

	//Decremeto +1
	x--
	fmt.Println("decremento de 1 de 11: ", x)

	//---------------------------------------------TIPOS DE DATOS PRIMITIVOS---------------------------------------------
	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

}
