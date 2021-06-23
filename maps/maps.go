package maps

import "log"

type User struct {
	FirstName string
	LastName  string
}

func Execute() {
	myMap := make(map[string]string) // keys  = string, values = string
	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"
	log.Println(myMap["dog"], myMap["other-dog"])

	myMap2 := make(map[string]int) // keys  = string, values = integers
	myMap2["first"] = 1
	myMap2["second"] = 2
	log.Println(myMap2["first"], myMap2["second"])

	myMap3 := make(map[string]User)
	vin := User{
		FirstName: "Vin",
		LastName:  "Diesel",
	}
	mare := User{
		FirstName: "Mare",
		LastName:  "Shioban",
	}
	myMap3["vin"] = vin
	myMap3["mare"] = mare
	log.Println(myMap3["vin"].LastName)
	log.Println(myMap3["mare"].LastName)

	// Ranging over maps
	log.Println(">>>>> RANGE OVER MAPS <<<<<<<")
	animals := make(map[string]string)
	animals["dog"] = "Sam"
	animals["bird"] = "Jimmy"
	for animalType, animal := range animals {
		log.Println(animalType, animal)
	}

}
