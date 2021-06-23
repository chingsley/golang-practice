package main

import (
	"testing"

	"github.com/chingsley/testing-go-things/testing_course"
)

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect 5", 100.0, 20.0, 5.0, false},
	{"expect fraction", 5.0, 2.0, 2.5, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := testing_course.Divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("got an error instead of a valid result. Error: ", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := testing_course.Divide(10.0, 1.0)
	if err != nil {
		t.Error(("Test failed, got an error for 10.0/1.0 instead of 10.0"))
	}
}
