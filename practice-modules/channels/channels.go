package channels

import (
	"fmt"

	"github.com/chingsley/testing-go-things/helpers"
)

func Execute() {

	// CREATING A CHANNEL
	integerChannel := make(chan int)
	defer close(integerChannel)

	go CalculateValue(integerChannel)

	num := <-integerChannel
	fmt.Println(num)

}

const numPool = 1000

func CalculateValue(intChannel chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChannel <- randomNumber
}
