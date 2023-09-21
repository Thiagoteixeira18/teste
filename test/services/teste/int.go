package services

import "fmt"

func SomeValue() string {
	return "expect value"
}

func Principal() {
	SomeValue := SomeValue()

	addition := Sum(3, 2)

	Multiplicacao := Multpy(5, -6)

	fmt.Printf("Some Value: %s, Addition: %d, multiplicacao: %d/n",
		SomeValue, addition, Multiplicacao)
}

func Sum(a int, b int) int {
	return a + b
}

func Multpy(a int, b int) int {
	return a * b
}

func multi(d, t int) int {
	return d * t
}


const quantidadeRepeticoes = 5

func Repetir(caractere string) string {
    var repeticoes string
    for i := 0; i < quantidadeRepeticoes; i++ {
        repeticoes += caractere
    }
    return repeticoes
}