package testing_course

import (
	"errors"
	"log"
)

func Divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}

	result = x / y
	return result, nil
}

func Execute() {
	const x = 100.0
	const y = 10.0
	result, err := Divide(x, y)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(x, "/", y, " = ", result)
}
