package tdd

import "testing"

func TestSayInput1Get1(t *testing.T) {
	s := Say(1)
	if s != "1" {
		t.Errorf("should be %v but got %v", "1", s)
	}
}

func TestSayInput2Get2(t *testing.T) {
	s := Say(2)
	if s != "2" {
		t.Errorf("should be %v but got %v", "2", s)
	}
}
