package main

import "fmt"

/* Calculado para a aula de GO da LinuxTips

Versão 0.0.1
Autor Italo Torres
Licença Unlisence

retorno da função main
func main() {
	fmt.Println(sum(1,2))
	fmt.Println(multiply(5,20))
	fmt.Println(divide(10,2))
	fmt.Println(substract(10, 10))
}
*/

func sum(a int, b int) int {
	return a + b
}
func multiply(a int, b int) int {
	return a * b
}
func divide(a float64, b float64) float64 {
	return a / b
}
func substract(a int, b int) int {
	return a - b
}

func main() {
	fmt.Println("O resultado da soma é:", sum(1, 2))
	fmt.Println("O resultando da multiplicação é:", multiply(5, 20))
	fmt.Println("O resultando da divisão é:", divide(10, 2))
	fmt.Println("O resultando da subtração é:", substract(10, 10))
}
