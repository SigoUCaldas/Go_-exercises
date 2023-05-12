package main

import (
	"fmt"
	"strconv"
)

func main() {
	// sumarTeclado()
	// sumar10Teclado()
	// sumarNTeclado()
	// galonesLitros()
	// galonesLitros100()
	// pesoLuna()
	// numeroCifras()
	// conversionPies()
	// pares100()
	// imparesN()
	// paresN()
	// multiplos3N()
	// fibonacci()
	// factura()
	promedioEstudiantes()

}
func sumarTeclado() {
	fmt.Print("Escriba numero 1: ")
	var numero1 int
	fmt.Scanln(&numero1)
	fmt.Print("Escriba numero 2: ")
	var numero2 int
	fmt.Scanln(&numero2)
	suma := numero1 + numero2
	fmt.Println("la suma es:", suma)

}
func sumar10Teclado() {
	var numero = 0
	suma := 0
	for i := 1; i <= 10; i++ {
		fmt.Print("Escriba numero ", i, ": ")
		fmt.Scanln(&numero)
		suma = suma + numero
	}
	fmt.Println("la suma es:", suma)
}
func sumarNTeclado() {
	count := 0
	fmt.Print("Escriba numero de elementos a sumar: ")
	fmt.Scanln(&count)
	var numero = 0
	suma := 0
	for i := 1; i <= count; i++ {
		fmt.Print("Escriba numero ", i, ": ")
		fmt.Scanln(&numero)
		suma = suma + numero
	}
	fmt.Println("la suma es:", suma)
}
func galonesLitros() {
	galones := 0.0
	fmt.Print("Escriba numero de galones: ")
	fmt.Scanln(&galones)
	litros := galones * 3.7854
	fmt.Println("la cantidad de litros son :", litros)
}
func galonesLitros100() {
	for i := 1; i <= 100; i++ {
		litros := float64(i) * 3.7854
		fmt.Println(i, "galones = ", litros, "litros")
	}
}

func pesoLuna() {
	peso := 0.0
	fmt.Print("Escriba su peso: ")
	fmt.Scanln(&peso)
	pesoLuna := peso * 0.17
	fmt.Println("Su peso efectivo en la luna es:", pesoLuna)
}
func numeroCifras() {
	numero := ""
	fmt.Print("Escriba el numero: ")
	fmt.Scanln(&numero)
	cifras := len(numero)
	fmt.Println("El numero de cifras es:", cifras)
}
func conversionPies() {
	pies := 0.0
	fmt.Print("Escriba numero de pies: ")
	fmt.Scanln(&pies)
	pulgadas := pies * 12
	cm := pulgadas * 2.54
	metros := cm / 100
	yardas := cm / 91.4

	fmt.Println("Pulgadas: ", pulgadas, "\nCentimetros: ", cm, "\nMetros: ", metros, "\nYardas: ", yardas)
}
func pares100() {
	numero := 1
	salida := 0
	for salida < 100 {
		if numero%2 == 0 {
			fmt.Println(numero)
			salida++
		}
		numero++
	}
}

func imparesN() {
	numero := 1
	salida := 0
	cantidad := 0
	fmt.Print("Escriba la cantidad de numeros: ")
	fmt.Scanln(&cantidad)
	for salida < cantidad {
		if numero%2 != 0 {
			fmt.Println(numero)
			salida++
		}
		numero++
	}
}
func paresN() {
	numero := 1
	salida := 0
	cantidad := 0
	fmt.Print("Escriba la cantidad de numeros: ")
	fmt.Scanln(&cantidad)
	for salida < cantidad {
		if numero%2 == 0 {
			fmt.Println(numero)
			salida++
		}
		numero++
	}
}
func multiplos3N() {
	numero := 1
	salida := 0
	cantidad := 0
	fmt.Print("Escriba la cantidad de numeros: ")
	fmt.Scanln(&cantidad)
	for salida < cantidad {
		if numero%3 == 0 {
			fmt.Println(numero)
			salida++
		}
		numero++
	}
}
func fibonacci() {
	a1, a2 := 1, 1
	enesimo := 0
	fmt.Print("Escriba el numero del termino: ")
	fmt.Scanln(&enesimo)
	salida := 1
	for i := 3; i <= enesimo; i++ {
		salida = a1 + a2
		a1 = a2
		a2 = salida
	}
	fmt.Print("El", enesimo, " termino es: ", salida)
}
func factura() {
	importe, iva := 0.0, 0.0
	fmt.Print("Escriba el importe:")
	fmt.Scanln(&importe)
	fmt.Print("Escriba el iva: ")
	fmt.Scanln(&iva)
	totalImporte := 0.0
	totalIva := 0.0
	for importe > 0 {
		if iva == 4 || iva == 7 || iva == 16 {
			totalImporte = +importe
			totalIva = +iva * importe / 100

		} else {
			fmt.Println("Error")
		}
		fmt.Print("Escriba el importe:")
		fmt.Scanln(&importe)
		fmt.Print("Escriba el iva: ")
		fmt.Scanln(&iva)
	}
	if totalImporte >= 1000 && totalImporte <= 10000 {
		totalImporte -= totalImporte * 5 / 100
		totalIva -= totalIva * 5 / 100
	} else if totalImporte > 10000 {
		totalImporte -= totalImporte * 10 / 100
		totalIva -= totalIva * 10 / 100
	}
	fmt.Println("Importe: ", totalImporte, "IVA: ", totalIva, "Total: ", totalImporte+totalIva)
}
func numerosN() {
	suma := 0
	producto := 1
	N := 0
	count := 1
	numero := 0
	fmt.Print("Escriba la cantidad de numeros: ")
	fmt.Scanln(&N)
	for count <= N {
		fmt.Print("Escriba el numero", count, ": ")
		fmt.Scanln(&numero)
		if numero%2 == 0 {
			suma += numero
		} else {
			producto *= numero
		}
		count++
	}
	fmt.Println("Suma:", suma, "   Producto:", producto)
}
func promedioEstudiantes() {
	numero := 0.0
	salida := ""
	for i := 1; i < 5; i++ {
		promedio := 0.0
		for j := 1; j < 3; j++ {
			fmt.Print("Ingrese nota ", j, " para el estudiante ", i, ": ")
			fmt.Scanln(&numero)
			promedio += numero
		}
		promedio /= 7
		salida = (salida + "\n  Promedio estudiante " + strconv.Itoa(i) + ": " + fmt.Sprintf("%f", promedio))
	}
	fmt.Print(salida)
}

// func gravity() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Println("Ingrese su peso:")
// 	scanner.Scan()
// 	valor, err := strconv.Atoi(scanner.Text())
// 	//var resultado = (float64(valor)) * 0.17
// 	var gravity = 9.8 * 0.17
// 	var resultado = ((float64(valor)) / 9.8) * gravity
// 	fmt.Println("Su peso en la Luna sera de: ", resultado, err)
// }

// go run ejercicios.go
