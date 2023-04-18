package main

import (
	"fmt"
	"math"
)

func inputValues() (int, int) {
	var input1 int
	var input2 int
	fmt.Print("Introduce el primer valor: ")
	fmt.Scan(&input1)
	fmt.Print("Introduce el segundo valor: ")
	fmt.Scan(&input2)
	return input1, input2
}

func operacion() {
	fmt.Println("Opciones disponibles:")
	fmt.Println("0) Salir")
	fmt.Println("1) Sumar")
	fmt.Println("2) Restar")
	fmt.Println("3) Multiplicar")
	fmt.Println("4) Dividir")
	fmt.Println("5) Obtener el módulo")
	fmt.Println("6) Cambiar de base")
	fmt.Println("7) Elevar un número")
	fmt.Println("8) Raíz cuadrada")
	fmt.Println("9) Área de un círculo")
	fmt.Println("10) Generar número PI con 3 decimales")

	var opcion int
	fmt.Print("Introduce una opcion: ")
	fmt.Scan(&opcion)
	switch opcion {
	case 0:
		fmt.Println("¡Se ha salido!")
		return
	case 1:
		input1, input2 := inputValues()
		fmt.Println("Resultado: ", input1+input2)
	case 2:
		input1, input2 := inputValues()
		fmt.Println("Resultado: ", input1-input2)
	case 3:
		input1, input2 := inputValues()
		fmt.Println("Resultado: ", input1*input2)
	case 4:
		input1, input2 := inputValues()
		fmt.Println("Resultado: ", input1/input2)
	case 5:
		input1, input2 := inputValues()
		fmt.Println("Resultado: ", input1%input2)
	case 6:
		var baseActual int
		var baseDestino int
		fmt.Print("Introduce el número: ")
		fmt.Scan(&baseActual)
		fmt.Print("Introduce la base de destino: ")
		fmt.Scan(&baseDestino)
		if baseDestino > 0 && baseDestino <= 16 {
			fmt.Println("Resultado: ", fmt.Sprintf("%x", baseActual))
		} else {
			fmt.Println("Valor introducido incorrecto")
		}
	case 7:
		var base float64
		var exponente float64
		fmt.Print("Introduce base: ")
		fmt.Scan(&base)
		fmt.Print("Introduce exp: ")
		fmt.Scan(&exponente)
		fmt.Println("Resultado: ", math.Pow(base, exponente))
	case 8:
		var numero float64
		fmt.Print("Introduce el número: ")
		fmt.Scan(&numero)
		fmt.Println("Resultado: ", math.Sqrt(numero))
	case 9:
		var radio float64
		fmt.Print("Introduce el radio del circulo: ")
		fmt.Scan(&radio)
		area := math.Pi * math.Pow(radio, 2)
		fmt.Println("El área del círculo es: ", area)
	case 10:
		fmt.Println("Valor de Pi: ", math.Pi)
	default:
		fmt.Println("Opción inválida")
	}
}

func main() {
	operacion()
}

// Este código es una implementación de una calculadora mejorada en el lenguaje de programación Go. Al ejecutar el programa, se muestra una lista de opciones disponibles para el usuario, que pueden incluir desde sumar, restar, multiplicar y dividir hasta calcular el módulo de dos números, cambiar de base, elevar un número a un exponente, calcular la raíz cuadrada, calcular el área de un círculo, y generar el número Pi con 3 decimales.

// La función "inputValues" permite al usuario introducir dos valores. La función "operacion" se encarga de realizar las operaciones matemáticas según la opción elegida por el usuario.

// El programa utiliza la biblioteca estándar "fmt" para leer y escribir valores en la consola y la biblioteca "math" para realizar operaciones matemáticas.
