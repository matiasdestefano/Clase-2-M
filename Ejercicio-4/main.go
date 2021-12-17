package main

import ("fmt"; "errors")

// Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso,
// requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

// Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
// y que devuelva otra función ( y un error en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros
// y devuelva el cálculo que se indicó en la función anterior


func main() {
	var nombreOperacion string = "promedio"
	var errorOperacion error

	operacion := realizarOperacion(nombreOperacion)
	resultado, errorOperacion := operacion(9, 7, 5, 10)

	if errorOperacion != nil {
		fmt.Printf("%s", errorOperacion)
	} else {
		fmt.Printf("El %s es de: %.2f", nombreOperacion, resultado)
	}	
}

func realizarOperacion(operacion string) func(notas ...int) (float32, error) {
	var errorNota error
	switch operacion {
	case "minimo":
		return func (notas ...int) (float32, error) {
			var minimo float32
			
			for i, nota := range notas {
				if i == 0 {
					minimo = float32(nota)
				}

				if float32(nota) < minimo {
					minimo = float32(nota)
				}

				if nota < 0 {
					errorNota = errors.New("Las notas ingresadas no son validas.")
					return 0, errorNota
				}
			}
			return float32(minimo), nil

		}
	case "maximo":
		return func (notas ...int) (float32, error) {
			var maximo float32
			
			for i, nota := range notas {
				if i == 0 {
					maximo = float32(nota)
				}

				if float32(nota) > maximo {
					maximo = float32(nota)
				}

				if nota < 0 {
					errorNota = errors.New("Las notas ingresadas no son validas.")
					return 0, errorNota
				}
			}
			return float32(maximo), nil
		}
	case "promedio":
		return func (notas ...int) (float32, error) {
			var promedio float32
			var sumaDeNotas float32
			var cantidadElementos = len(notas)
			
			for _, nota := range notas {
				if nota < 0 {
					return 0, errors.New("Las notas ingresadas no son validas")
				}
				sumaDeNotas += float32(nota)
			}

			promedio = sumaDeNotas / float32(cantidadElementos)
			return promedio, nil
		}
	default:
		return nil
	}
}