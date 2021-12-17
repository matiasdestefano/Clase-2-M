package main

import ("fmt"; "errors")

//Una empresa marinera necesita calcular el salario de sus empleados basándose en la
//cantidad de horas trabajadas por mes y la categoría.

//Si es categoría C, su salario es de $1.000 por hora
//Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
//Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

//Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

func main() {
	var categoria string = "A"
	var horasTrabajadas int = 160
	var errorSalario error

	resultado, errorSalario := calcularSalario(horasTrabajadas, categoria)

	if(errorSalario != nil) {
		fmt.Printf("%s", errorSalario)
	} else {
		fmt.Printf("El salario del empleado de categoria %s es de: $%.2f", categoria,resultado)
	}
}

func calcularSalario(horasTrabajadas int, categoria string) (float32, error) {

	var bonoMensual float32
	var valorHora float32
	var errorCategoria error

	bonoMensual, errorCategoria = calcularBono(categoria)
	valorHora, errorCategoria = calcularValorHora(categoria)

	if errorCategoria != nil {
		return 0, errorCategoria
	}

	if(horasTrabajadas < 0) {
		return 0, errors.New("La cantidad de horas trabajadas no es valida.")
	}
	salarioTotal := float32(horasTrabajadas) * valorHora
	return salarioTotal + salarioTotal * bonoMensual, nil

}

func calcularBono(categoria string) (float32, error) {
	if categoria == "A" {
		return 0.5, nil
	} else if categoria == "B" {
		return 0.2, nil
	} else if categoria == "C" {
		return 0, nil
	} else {
		return 0, errors.New("La categoria no es valida.")
	}
}

func calcularValorHora(categoria string) (float32, error) {
	if categoria == "A" {
		return 3000, nil
	} else if categoria == "B" {
		return 1500, nil
	} else if categoria == "C" {
		return 1000, nil
	} else {
		return 0, errors.New("La categoria no es valida.")
	}
}