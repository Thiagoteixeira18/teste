package services_test

import (
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

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("'Mundo' como padrão para 'string' vazia", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris", "")
		esperado := "Olá, Chris"
		verificaMensagemCorreta(t, resultado, esperado)
	})
	
	t.Run("bonjur", func(t *testing.T) {
		resultado := Ola("", frances)
		esperado := "Bonjour, Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}