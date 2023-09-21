package services_test

import "fmt"

func Principal() {
	SomeValue := SomeValue()

	addition := Sum(3, 2)

	Multiplicacao := Multpy(5, -6)

	fmt.Printf("Some Value: %s, Addition: %d, multiplicacao: %d/n",
		SomeValue, addition, Multiplicacao)
}

func SomeValue() string {
	return "expect value"
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

const (
    espanhol = "espanhol"
    frances  = "frances"
    prefixoOlaPortugues = "Olá, "
    prefixoOlaEspanhol  = "Hola, "
    prefixoOlaFrances   = "Bonjour, "
)

func Ola(nome string, idioma string) string {
    if nome == "" {
        nome = "Mundo"
    }

    return prefixodeSaudacao(idioma) + nome
}

func prefixodeSaudacao(idioma string) (prefixo string) {
    switch idioma {
    case frances:
        prefixo = prefixoOlaFrances
    case espanhol:
        prefixo = prefixoOlaEspanhol
    default:
        prefixo = prefixoOlaPortugues
    }
    return
}