package main

import ("fmt"; "errors")

//Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. 
//Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan
//haber muchos más animales que refugiar.

//Por perro necesitan 10 kg de alimento
//Por gato 5 kg
//Por cada Hamster 250 gramos.
//Por Tarántula 150 gramos.


//Se solicita:
//Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado
//y que retorne una función y un error (en caso que no exista el animal)
//Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

const (
	perro = "perro"
	gato = "gato"
	hamster = "hamster"
	tarantula = "tarantula"
)

func main() {
	
	animalPerro, errorAnimal := Animal(perro)
	animalGato, errorAnimal := Animal(gato)
	animalHamster, errorAnimal := Animal(hamster)
	animalTarantula, errorAnimal := Animal(tarantula)

	if errorAnimal == nil {
		fmt.Printf("El alimento para %s sera de: %.2f Kg\n", perro, animalPerro(3))
		fmt.Printf("El alimento para %s sera de: %.2f Kg\n", gato, animalGato(10))
		fmt.Printf("El alimento para %s sera de: %.2f Kg\n", hamster, animalHamster(5))
		fmt.Printf("El alimento para %s sera de: %.2f Kg\n", tarantula, animalTarantula(1))

	}
}

func Animal(animal string) (func (cantidad int) float32, error) {

	const (
		alimentoPerro = 10.0
		alimentoGato = 5.0
		alimentoHamster = 0.25
		alimentoTarantula = 0.15
	)

	switch animal {
	case perro:
		return func(cantidad int) float32 {
			return (float32(cantidad) * float32(alimentoPerro))
		}, nil
		case gato:
			return func(cantidad int) float32 {
				return (float32(cantidad) * float32(alimentoGato))
			}, nil
		case hamster:
			return func(cantidad int) float32 {
				return (float32(cantidad) * float32(alimentoHamster))
			}, nil
		case tarantula:
			return func(cantidad int) float32 {
				return (float32(cantidad) * float32(alimentoTarantula))
			}, nil
		default:
			return nil, errors.New("El animal no existe.")
	}
}