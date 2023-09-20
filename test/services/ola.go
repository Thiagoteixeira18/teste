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
