package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
)

func main() {
	var name string = "Jeff"
	surname, age := "Carnass", 25
	const basicPlanName = "Basic Plan"

	fmt.Println(name, surname, age, basicPlanName)
	fmt.Println("My favorite number is", rand.Intn(10))

	myNewNumbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	myNewNumbers = append(myNewNumbers, 10)
	mySlice := make([]int, 5)
	mySlice2 := []int{}
	fmt.Println(myNewNumbers, mySlice, mySlice2)

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
	mynewNumber := aggregate(3, 4, 5, mul)
	fmt.Println(aggregate(2, 3, 4, mul))
	fmt.Println("my new number is ", mynewNumber)
	fmt.Println("hello there new")

	fmt.Println(Micheal)

	weatherReport := concatter()
	weatherReport("Todays weather")
	weatherReport("extremeley hot")

	fmt.Println(weatherReport("please be safe"))

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
func sliceRange() {
	fruits := []string{"apple", "banana", "grape"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
}
func mapsStart() {
	ages := map[string]int{
		"John": 37,
		"Mary": 21,
	}
<<<<<<< HEAD
	fmt.Println(len(ages)) // 2	
=======
	fmt.Println(len(ages)) // 2
}
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

func mul(x, y int) int {
	return x * y
}

// CopyFile copies a file from srcName to dstName on the local filesystem.
func CopyFile(dstName, srcName string) (written int64, err error) {

	// Open the source file
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	// Close the source file when the CopyFile function returns
	defer src.Close()

	// Create the destination file
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	// Close the destination file when the CopyFile function returns
	defer dst.Close()

	return io.Copy(dst, src)
}
func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
>>>>>>> 3bf2291360b3ce33d68b7f4da9b29a99a09eee89
}
