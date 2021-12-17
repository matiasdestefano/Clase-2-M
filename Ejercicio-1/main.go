package main

import "fmt"

func main() {
	var salario, impuesto float32
	salario = 200000

	impuesto = calcularImpuesto(salario)
	fmt.Printf("El impuesto es de $%.2f", impuesto)
}

func calcularImpuesto(salario float32) float32 {

	var impuestoMayor float32 = 0.27
	var impuestoMenor float32 = 0.17

	if salario > 150000 {
		return salario * impuestoMayor
	} else if salario > 50000 {
		return salario * impuestoMenor
	} else {
		return 0
	}
}