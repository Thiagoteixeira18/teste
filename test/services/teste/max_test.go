package services

import "testing"

func TestMax(t *testing.T) {
	max := Max(5, 7)
	if max != 7 {
		t.Errorf("Max incorreto, esperava 7, obteve %d", max)
	}
}
