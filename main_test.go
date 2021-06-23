package main

import (
	"testing"

	"github.com/chingsley/testing-go-things/testing_course"
)

func TestDivision(t *testing.T) {
	_, err := testing_course.Divide(10.0, 1.0)
	if err != nil {
		t.Error(("Test failed, got an error for 10.0/1.0 instead of 10.0"))
	}
}
