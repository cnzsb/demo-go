package main

import "fmt"

type Animal struct {
	Legs int
	Age  int

	// private
	name  string
	power string
}

func (animal Animal) Info() {
	fmt.Println("Animal Info: ", animal.Legs, animal.Age, animal.name, animal.power)
}

// private
func (animal Animal) setName(name string) {
	animal.name = name
}

// not recommendation as can directly update the value
func (animal *Animal) updateName(name string) {
	// equal as animal
	// (*animal).name = name
	animal.name = name
}

// SetAnimalPower will not modify the original value
func SetAnimalPower(animal Animal, power string) {
	animal.power = power
}

// UpdateAnimalPower needing a pointer
func UpdateAnimalPower(animal *Animal, power string) {
	animal.power = power
}

// Cat inherit
type Cat struct {
	Animal

	tail bool
}

// if name it as Info, will reset the Animal.Info
func (cat *Cat) info() {
	fmt.Println("Cat Info: ", cat.Legs, cat.Age, cat.name, cat.power, cat.tail)
}

func main() {
	dog1 := Animal{Legs: 2, Age: 2}
	dog1.name = "dog1"
	dog1.Info()

	dog2 := Animal{4, 5, "dog2", ""}
	dog2.Info()
	SetAnimalPower(dog2, "Guard")
	dog2.Info()
	UpdateAnimalPower(&dog2, "Guard")
	dog2.Info()
	// update directly
	dog2.Age = 10
	dog2.Info()

	dog2.setName("DOG2")
	dog2.Info()
	dog2.updateName("DOG2")
	dog2.Info()

	cat1 := Cat{Animal: Animal{Legs: 4, Age: 1}, tail: true}
	cat1.Info()

	cat2 := Cat{Animal{Legs: 4, Age: 2}, false}
	cat2.Info()
	cat2.info()

	var cat3 Cat
	cat3.name = "cat3"
	cat3.Legs = 4
	cat3.tail = true
	cat3.info()
}
