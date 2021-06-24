// +build ignore
package slices

import (
	"log"
	"sort"
)

func Execute() {
	var mySlice []string
	mySlice = append(mySlice, "vin Diesel")
	mySlice = append(mySlice, "Mare Sheehan")
	log.Println((mySlice))

	var mySlice2 []int
	mySlice2 = append(mySlice2, 3)
	mySlice2 = append(mySlice2, 2)
	mySlice2 = append(mySlice2, 1)
	log.Println((mySlice2))
	sort.Ints(mySlice2)
	log.Println((mySlice2))

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	log.Println(numbers)
	log.Println(numbers[0:2])
	log.Println(numbers[3:7])

	names := []string{"Mare", "Lori", "Soibhan", "Zabel", "Ryan", "Dylan", "McMenamin"}
	log.Println(names)
	log.Println(names[0:2])
	log.Println(names[3:7])

}
