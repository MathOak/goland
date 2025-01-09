package main

import (
	"fmt"
	"os"
	"os/exec"
)

func add(n1, n2 int) int {
	return n1 + n2
}
func getAddNumbers() {
	var n1, n2 int
	fmt.Print("Entre com o primeiro e segundo número:")
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Printf("\nO valor somado de %d + %d é %d", n1, n2, add(n1, n2))
}

func sub(n1, n2 int) int {
	return n1 - n2
}
func getSubNumbers() {
	var n1, n2 int
	fmt.Print("Entre com o subtraendo e o minuendo:")
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Printf("\nO valor subtraido de %d - %d é %d", n1, n2, sub(n1, n2))
}

func mult(n1, n2 int) int {
	return n1 * n2
}
func getMultNumbers() {
	var n1, n2 int
	fmt.Print("Entre com os fatores:")
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Printf("\nO produto de %d * %d é %d", n1, n2, mult(n1, n2))
}

func div(n1, n2 int) (float64, error) {
	if n2 == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return float64(n1) / float64(n2), nil
}
func getDivNumbers() {
	var n1, n2 int
	fmt.Print("Entre com o dividendo e o divisor:")
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	var result, err = div(n1, n2)
	if err != nil {
		fmt.Printf("\nO quociente de %d / %d é %g", n1, n2, result)
	} else {
		fmt.Println(err)
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clearScreen()
	var inputOperation string
	fmt.Println("Calculadora\nOperações disponiveis: soma, subtração, multiplicação, divisão\nUtilize os seguintes comandos: 'sum'ou '+', 'sub' ou '-', 'mult' ou '*', 'div' ou '/'\nDigite 'exit' para sair")
	fmt.Print("Entre com a operação desejada: ")
	fmt.Scanf("%s", &inputOperation)

	switch inputOperation {
	case "sum":
		getAddNumbers()
	case "+":
		getAddNumbers()
	case "sub":
		getSubNumbers()
	case "-":
		getSubNumbers()
	case "mult":
		getMultNumbers()
	case "*":
		getMultNumbers()
	case "div":
		getDivNumbers()
	case "/":
		getDivNumbers()

	case "exit":
		fmt.Println("\nSaindo...")
	default:
		fmt.Println("\nOperação inválida")
		fmt.Println("\nSaindo...")

	}

}
