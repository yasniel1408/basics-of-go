package main

import "fmt"

func printHollaMundo(num string) {
	fmt.Println("Hola Mundo ", num)
}

// aqui puede ser => (a int, b int, c string) ó (a,b int, c string)
func tripeArgument(a, b int, c string) {
	fmt.Println("Tres datos ", a, b, c)
}

func returnValue(a int) int {
	value := a * 2
	fmt.Println("return value ", value)
	return value
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	printHollaMundo("1")
	printHollaMundo("2")

	tripeArgument(1, 2, "tres")

	returnValue(2)

	value1, value2 := doubleReturn(2)
	fmt.Println("double return ", value1, value2)

	// si solo quiero un valor
	value11, _ := doubleReturn(2)
	fmt.Println("double return ", value11)
}
