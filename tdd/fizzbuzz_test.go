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

func TestSayInput3GetBizz(t *testing.T) {
	s := Say(3)
	if s != "Fizz" {
		t.Errorf("should be %v but got %v", "Fizz", s)
	}
}

func TestSayInput4Get4(t *testing.T) {
	s := Say(4)
	if s != "4" {
		t.Errorf("should be %v but got %v", "4", s)
	}
}

func TestSayInput5GetBuzz(t *testing.T) {
	s := Say(5)
	if s != "Buzz" {
		t.Errorf("should be %v but got %v", "Buzz", s)
	}
}

func TestSayInput6GetFizz(t *testing.T) {
	s := Say(6)
	if s != "Fizz" {
		t.Errorf("should be %v but got %v", "Fizz", s)
	}
}

func TestSayInput7Get7(t *testing.T) {
	s := Say(7)
	if s != "7" {
		t.Errorf("should be %v but got %v", "7", s)
	}
}

func TestSayInput9GetFizz(t *testing.T) {
	s := Say(9)
	if s != "Fizz" {
		t.Errorf("should be %v but got %v", "Fizz", s)
	}
}

func TestSayInput10GetBuzz(t *testing.T) {
	s := Say(10)
	if s != "Buzz" {
		t.Errorf("should be %v but got %v", "Buzz", s)
	}
}

func TestSayInput12GetFizz(t *testing.T) {
	s := Say(12)
	if s != "Fizz" {
		t.Errorf("should be %v but got %v", "Fizz", s)
	}
}

func TestSayInput15GetFizzBuzz(t *testing.T) {
	s := Say(15)
	if s != "FizzBuzz" {
		t.Errorf("should be %v but got %v", "FizzBuzz", s)
	}
}
