package main

import (
	"fmt"

	"github.com/anara123/hellogo/stringutil"
)

func main() {
	packages()
	variables()
	closures()
	dynamicTyping()
	constants()
	memoryAddress()
	pointers()
	conditionBreakContinue()
	runes()
	switchStatement()
	ifStatement()
	variadicFunctions()
}

func pointers() {
	fmt.Println("Pointers\n====================")
	zero := func(z int) {
		fmt.Println(&z)
		z = 0
	}

	x := 5
	fmt.Println(x)
	fmt.Println(&x)
	zero(x)
	fmt.Println(x)

	// =======
	fmt.Println("=============")
	zeroPointer := func(z *int) {
		fmt.Println(z)
		*z = 0
	}
	fmt.Println(x)
	fmt.Println(&x)
	zeroPointer(&x)
	fmt.Println(x)
}

func memoryAddress() {
	fmt.Println("Memory Address\n====================")
	a := 5

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", &a)

	// var meters float64
	// fmt.Println("How many meters do you sweem?")
	// fmt.Scan(&meters)

	// fmt.Printf("You are swimming %f meters\n", meters)
}

const p string = "this is constant"

const (
	_  = iota // 0
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func constants() {
	fmt.Println("Constants\n====================")
	const q = "this is inner const"

	fmt.Printf("%T \t%v\n", p, p)
	fmt.Printf("%T \t%v\n", q, q)

	fmt.Printf("KB %v\n", KB)
	fmt.Printf("MB %v\n", MB)
	fmt.Printf("GB %v\n", GB)
	fmt.Printf("TB %v\n", TB)
}

func packages() {
	fmt.Println("Packages\n====================")

	s := stringutil.Reverse("hello")
	fmt.Println(s)
}

func variables() {
	fmt.Println("Variables\n====================")

	max := max(7)
	fmt.Printf("%v", max)
	fmt.Println()
}

func closures() {
	fmt.Println("Closures\n====================")

	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())

	user := UserCreator("Anar")("Azadaliyev")(29)("123qwe")
	fmt.Println(user)
}

func dynamicTyping() {
	fmt.Println("Dynamic Typing\n====================")
	fmt.Printf("%v\t %v\t %v\n", conditinalReturn(11), conditinalReturn(7), conditinalReturn(9))
}

func conditionBreakContinue() {
	i := 0
	for {
		i++

		if i%2 == 0 {
			fmt.Println(i)
			continue
		}

		if i > 10 {
			break
		}
	}
}

func runes() {
	fmt.Println("runes")

	var a rune
	a = 'a'
	fmt.Println(rune(a))

	for i := 5000; i < 5020; i++ {
		fmt.Printf("%v\t- %v\t-%v\n", i, string(i), []byte(string(i)))
	}
}

// Contact is simple type for switch example
type Contact struct {
	fullname string
	phone    string
}

func switchStatement() {
	fmt.Println("switch statement")

	// switch with conditions
	myName := "koo"
	switch {
	case len(myName) == 2:
		fmt.Println("Your name length is 2")
	case myName == "koo":
		fmt.Println("Your name is koo")
	case myName == "moo":
		fmt.Println("Your name is moo")
	}

	// switch with fallthrough example
	switch "Anar" {
	case "Azer":
		fmt.Println("Hello Azer")
		fallthrough // if it will be Azer then it will fallthrough to Orxan
	case "Orxan":
		fmt.Println("Hello Orxan")
	case "Anar":
		fmt.Println("Hello Anar")
	case "Rufet":
		fmt.Println("Hello Rufet")
	default:
		fmt.Println("This is default")
	}

	switchByType(Contact{
		fullname: "Anar Azadaliyev",
		phone:    "0509957803",
	})
}

func switchByType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("This is an Integer")
	case string:
		fmt.Println("This is an String")
	case Contact:
		fmt.Println("This is an Contact")
	default:
		fmt.Println("There is no such type")
	}
}

func ifStatement() {
	if m := max(7); m <= 49 {
		fmt.Println("Less than 49")
	} else {
		fmt.Println("Bigger than 49")
	}
}

func variadicFunctions() {
	fmt.Println("Variadic Functions")
	fmt.Println(average(2, 4, 6, 10))

	fmt.Println("Variadic parameters")
	data := []float64{2, 4, 6, 10}
	fmt.Println(average(data...))
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)

	var total float64
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}

// ===========================================
//
//
// ===========================================

func max(x int) int {
	return 42 + x
}

func conditinalReturn(x int) interface{} {
	if x > 10 {
		return "str"
	} else if x < 8 {
		return 5
	} else {
		return 1.2
	}
}

// closures
// returns func() int - which is the type of the inner function
func wrapper() func() int {
	var x int

	return func() int {
		x++
		return x
	}
}

// UserCreator is curring function which uses closures to constrcut a User
func UserCreator(fname string) func(lname string) func(age int) func(password string) User {
	return func(lname string) func(age int) func(password string) User {
		return func(age int) func(password string) User {
			return func(password string) User {

				return User{
					fname:    fname,
					lname:    lname,
					age:      age,
					password: password,
				}
			}
		}
	}
}

// User is struct
type User struct {
	fname    string
	lname    string
	age      int
	password string
}
