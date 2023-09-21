package services

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	if Sum(2, 2) != 4 {
		t.Errorf("2 + 2 must be 4")
	}
}

func TestMultply(t *testing.T) {
	require.Equal(t, Multpy(2, 2), 4)
}

func TestSomeValue(t *testing.T) {
	result := SomeValue()
	if result != "expect value" {
		t.Errorf("SomeValue() = %q, want %q", result, "expect value")
	}
}

func TestMulti(t *testing.T) {
	tests := []struct {
		d, t, esperado int
	}{
		{2, 3, 6},
		{-2, 3, -6},
		{2, -3, -6},
		{0, 3, 0},
	}
	for _, test := range tests {
		result := multi(test.d, test.t)
		if result != test.esperado {
			t.Errorf("Multi(%d, %d) = %d; want %d", test.d, test.t, result, test.esperado)
		}
	}
}

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a")
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}
}
func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a")
	}
}

func TestSoma(t *testing.T) {

	t.Run("coleção de 5 números", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}

		resultado := Soma(numeros)
		esperado := 15

		if resultado != esperado {
			t.Errorf("resultado %d, want %d, dado %v", resultado, esperado, numeros)
		}
	})

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}

		resultado := Soma(numeros)
		esperado := 6

		if resultado != esperado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})

}

func TestSomaTodoOResto(t *testing.T) {

    verificaSomas := func(t *testing.T, resultado, esperado []int) {
        t.Helper()
        if !reflect.DeepEqual(resultado, esperado) {
            t.Errorf("resultado %v, esperado %v", resultado, esperado)
        }
    }

    t.Run("faz a soma do resto", func(t *testing.T) {
        resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
        esperado := []int{2, 9}
        verificaSomas(t, resultado, esperado)
    })

    t.Run("soma slices vazios de forma segura", func(t *testing.T) {
        resultado := SomaTodoOResto([]int{}, []int{3, 4, 5})
        esperado := []int{0, 9}
        verificaSomas(t, resultado, esperado)
    })

}