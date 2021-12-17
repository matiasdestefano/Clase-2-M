package main

import ("fmt"; "errors")

//Calcular promedio, recibir N enteros. Devolver error si existe un número negativo.
func main() {
	var promedio float32
	var errorPromedio error

	promedio, errorPromedio = calcularPromedio(8,-9,7,9,9,10)
	if errorPromedio != nil {
		fmt.Printf("%s", errorPromedio)
	} else {
		fmt.Printf("El promedio es de: %.2f", promedio)
	}
}

func calcularPromedio(notas ...float32) (float32, error) {
	cantidadDeNotas := float32(len(notas))
	var sumaDeNotas float32
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("Se ha ingresado una nota no válida.")
		}
		sumaDeNotas += nota
	}

	return (sumaDeNotas / cantidadDeNotas), nil

}