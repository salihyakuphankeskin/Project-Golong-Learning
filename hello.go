package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var name string = "Jeff"
	surname, age := "Carnass", 25
	const basicPlanName = "Basic Plan"

	fmt.Println(name, surname, age, basicPlanName)
	fmt.Println("My favorite number is", rand.Intn(10))

	type human struct {
		Name    string
		Surname string
		Age     int
	}
	type MedicalInfo struct {
		human
		bloodType string
	}
	type Rect struct {
		len, wid int
	}
	Micheal := human{"Micheal", "Hersey", 35}

	fmt.Println(Micheal)

	fmt.Scanln()
}

func typeFinder(value any) {
	fmt.Printf("The type of penniesPerText is %T\n", value)
}
func twoValueSameLine() {
	monthName, nextMonthName := "July", "August"
	fmt.Println(monthName, nextMonthName)
}
func floatToInt(value float64) int {

	valueInt := int(value)
	return valueInt
}
func secondsInHour() {
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	fmt.Println("number of seconds in an hour:", secondsInHour)
}
func printFunctions() {
	var myString string
	//	Interpolate the default representation
	myString = fmt.Sprintf("I am %v years old", 10)
	println(myString) // I am 10 years old

	myString = fmt.Sprintf("I am %v years old", "way too many")
	println(myString) // I am way too many years old

	myString = fmt.Sprintf("I am %s years old", "way too many")
	println(myString) // I am way too many years old

	myString = fmt.Sprintf("I am %d years old", 10)
	println(myString) // I am 10 years old

	myString = fmt.Sprintf("I am %f years old", 10.523)
	println(myString) // I am 10.523000 years old

	myString = fmt.Sprintf("I am %.2f years old", 10.523)
	println(myString) // I am 10.52 years old

	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %v, your open rate is %v percent", name, openRate)
	fmt.Println(msg)
}
func ifStatements() {
	if 1 > 0 {
		fmt.Println("one is bigger than zero")
	}
	appleColor := "Red"

	if currentColor := "Green"; currentColor == appleColor {
		fmt.Println("current color and apple color is matched")
	}
}
func add(x, y int) int {
	return x + y
}
func add2Times(f func(int, int) int, x int) int {
	result := f(x, x)
	return f(result, x)
}
func JohnAndDoe() (string, string) {
	return "John", "Doe"
}
func getJustJoe(x string) {
	x, _ = JohnAndDoe()
	fmt.Println(x)
}

// taking a struct
type Rect struct {
	len, wid int
}

func (re Rect) Area() int {
	return re.len * re.wid
}

type tank interface {
	// function categorizer
	Tarea() float64
}
type myvalue struct {
	radius float64
	height float64
}

func (m myvalue) Tarea() float64 {

	return 2*m.radius*m.height + 2*3.14*m.radius*m.radius
}

func sums(nums ...int) int {
	var num int = 0
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}
	return num

}
