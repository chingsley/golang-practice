package http_module

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "Hello, world!, this is the Home page")
	if err != nil {
		fmt.Println(err)
	}
	_, _ = fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page, 2 + 2 = %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x float32 = 100.0
	var y float32 = 0.0
	f, err := divideValue(x, y)
	fmt.Println(">>>>>>>>>>>\nerr = ", err)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("%s", err))
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f / % f = %f", x, y, f))
}

func divideValue(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("division by zero is not allowed")
		return 0.0, err
	}
	return x / y, nil
}

// addValues adds two integers
func addValues(x, y int) int {
	return x + y
}

func Execute() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println("Starting the applicatin on port: ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
