package main

import "fmt"

//doc -> https://golang.org/pkg/fmt/

func main() {
	helloMessage := "Hello"
	wordMessage := "Word"

	//Salto de linea de fmt
	fmt.Println(helloMessage, wordMessage)
	fmt.Println(helloMessage, wordMessage)

	//Printf // %s => string, %d => decimal, %f => float, %v => cualquier tipo de datos
	nombre := "Platzi"
	cursos := 500
	si := true
	fmt.Printf("%s tiene mas de %d cursos, dime que es %v \n", nombre, cursos, si)

	//Sprintf, guardar en una variable, puede ser usada para crear strings complejos
	message := fmt.Sprintf("%s tiene mas de %d cursos, dime que es %v \n", nombre, cursos, si)
	fmt.Printf(message)

	//Conocer tipo de datos puede servir para debugear
	fmt.Printf("helloMessage: %T \n", helloMessage)
	fmt.Printf("cursos: %T", cursos)
}
