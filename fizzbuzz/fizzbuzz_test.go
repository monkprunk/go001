package fizzbuzz

import "testing"

func TestSayInput1Return1(t *testing.T) {
	s := Say(1)
	if s != "1" {
		t.Errorf("should return %v but got %v ", "1", s)
	}
}

func TestSayInput2Return2(t *testing.T) {
	s := Say(2)
	if s != "2" {
		t.Errorf("should return %v but got %v ", "2", s)
	}
}

func TestSayInput3ReturnFizz(t *testing.T) {
	s := Say(3)
	if s != "Fizz" {
		t.Errorf("should return %v but got %v ", "Fizz", s)
	}
}

func TestSayInput5ReturnBuzz(t *testing.T) {
	s := Say(5)
	if s != "Buzz" {
		t.Errorf("should return %v but got %v ", "Buzz", s)
	}
}

func TestSayInput9ReturnFizz(t *testing.T) {
	s := Say(9)
	if s != "Fizz" {
		t.Errorf("should return %v but got %v ", "Fizz", s)
	}
}
func TestSayInput13Return13(t *testing.T) {
	s := Say(13)
	if s != "13" {
		t.Errorf("should return %v but got %v ", "13", s)
	}
}
func TestSayInput14Return14(t *testing.T) {
	s := Say(14)
	if s != "14" {
		t.Errorf("should return %v but got %v ", "14", s)
	}
}

func TestSayInput15ReturnFizzBuzz(t *testing.T) {
	s := Say(15)
	if s != "FizzBuzz" {
		t.Errorf("should return %v but got %v ", "FizzBuzz", s)
	}
}
