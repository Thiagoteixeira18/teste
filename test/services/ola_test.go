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


func TestMult(t *testing.T) {
 require.Equal(t, Multpy(2, 2), 4)
}