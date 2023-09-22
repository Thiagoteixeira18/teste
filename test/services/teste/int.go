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

func Soma(numeros []int) int {
    soma := 0
    for _, numero := range numeros {
        soma += numero
    }
    return soma
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
    var somas []int
    for _, numeros := range numerosParaSomar {
        if len(numeros) == 0 {
            somas = append(somas, 0)
        } else {
            final := numeros[1:]
            somas = append(somas, Soma(final))
        }
    }

    return somas
}

func Perimetro(largura float64, altura float64) float64 {
    return 2 * (largura + altura)
}

func Area(largura float64, altura float64) float64 {
    return largura * altura
}