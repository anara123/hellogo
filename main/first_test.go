package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	if average(1, 2, 8, 1) != 3 {
		t.Fail()
	}
}

func ExampleCallbacks() {
	visit([]int{55, 44, 33, 22}, func(num int) {
		fmt.Println(num)
	})

	// Output:
	// 55
	// 44
	// 33
	// 22
}

func TestHalf(t *testing.T) {
	assert := assert.New(t)

	result, isEven := half(5)
	assert.Equal(result, 2.5)
	assert.Equal(isEven, false)
}

func half(num int) (float64, bool) {
	return float64(num) / 2, num%2 == 0
}

type Student struct {
	name string
	age  int
}

func TestReferences(t *testing.T) {
	assert := assert.New(t)

	makeArr := make([]int, 3, 3)
	makeArr[0] = 1
	makeArr[1] = 2
	makeArr[2] = 3

	arr := [3]int{1, 2, 3}
	assert.NotEqual(arr, makeArr, "make and array are not equal")

	func(k []int) {
		assert.Equal(makeArr, k, "reference must be equal")
		assert.Equal(&makeArr[2], &k[2], "reference must be equal")
		if &makeArr[2] != &k[2] {
			fmt.Printf("item in makeArr references must be equal")
			t.Fail()
		}
	}(makeArr)

	func(m [3]int) {
		assert.Equal(arr, m, "reference must be equal")
		if &arr[2] == &m[2] {
			fmt.Printf("item in array when passed to function must not be equal")
			t.Fail()
		}

		m[2] = 55
		assert.NotEqual(arr[2], m[2], "The values must be not equal")
	}(arr)

	num := 10
	func(a int) {
		assert.Equal(num, a, "values must be equal")
		if &num == &a {
			fmt.Printf("%v\t%v references must be not equal", &num, &a)
			t.Fail()
		}
	}(num)

	koo(num, &num, assert)

	s := Student{"Anar", 29}
	func(otherS Student, otherSPointer *Student) {
		assert.Equal(s, otherS, "Student values must the same")

		otherS.name = "Fuad"
		assert.NotEqual(s.name, otherS.name, "names should be equal")

		assert.Equal(&s, otherSPointer, "Student values must the same")
		otherSPointer.name = "Fuad"
		assert.Equal(s.name, otherSPointer.name, "names should change")
	}(s, &s)
}

func koo(a int, expected *int, assert *assert.Assertions) {
	assert.Equal(*expected, a, "values must be equal")
	assert.Equal(expected, &a, "references must not be equal")

	var p1 *int
	var p2 *int

	k1 := 44
	k2 := 44
	p1 = &k1
	p2 = &k2
	assert.Equal(p1, p2, "k1 must be equal to k2")
}

func TestFilter(t *testing.T) {
	assert := assert.New(t)

	filtered := filter([]int{1, 2, 3, 4}, func(num int) bool {
		// is even
		return num%2 == 0
	})

	assert.Equal([]int{2, 4}, filtered, "Filtered not well")
}

// There are 3 ways to create slice.
func TestSlices(t *testing.T) {
	assert := assert.New(t)

	// using var
	var slice1 []int
	assert.Equal(true, slice1 == nil, "deafult value when using var is nil")
	assert.NotEqual(nil, slice1, "deafult value when using var is nil")
	assert.Equal([]int(nil), slice1, "deafult value when using var is nil")

	// immidiate initialization
	slice2 := []int{}
	assert.NotEqual([]int(nil), slice2)
	assert.Equal(true, slice2 != nil)
	assert.Equal(0, len(slice2))

	// using make
	slice3 := make([]int, 1)
	assert.Equal(1, len(slice3))
	assert.Equal(1, cap(slice3))
}

func TestDeleteFromSlice(t *testing.T) {
	slice := Sequence{22, 33, 44, 55}
	// delete 2nd element
	newSlice := append(slice[:2], slice[3:]...)

	assert.Equal(t, Sequence{22, 33, 55}, newSlice)

	assert.Equal(t, Sequence{22, 33, 55, 66}, slice.Delete(2).Push(66))
}

type Sequence []int

func (slice Sequence) Delete(index int) Sequence {
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

func (slice Sequence) Push(item int) Sequence {
	slice = append(slice, item)
	return slice
}

func ExampleConvertArrayToSlice() {
	arr := [4]int{1, 2, 3, 4}
	func(slice []int) {
		fmt.Println(slice)
	}(arr[:])

	// Output: [1 2 3 4]
}

func ExampleAnonymousFunc() {
	func() {
		fmt.Println("This is anonymous function")
	}()

	// Output: This is anonymous function
}

// defer executes the function before exiting the main function
// similar somehow to finially
func ExampleDefer() {
	defer world()
	hello()

	// Output: Hello World
}

func world() {
	fmt.Print("World")
}

func hello() {
	fmt.Print("Hello ")
}

func filter(arr []int, condition func(int) bool) []int {
	var result = []int{}
	for _, v := range arr {
		if condition(v) {
			result = append(result, v)
		}
	}

	return result
}

func visit(numbers []int, callback func(int)) {
	for _, num := range numbers {
		callback(num)
	}
}

func ExamplePrintf() {
	fmt.Printf("Hello %v", "Anar")
	// Output: Hello Anar
}

func TestMaps(t *testing.T) {
	var assert = assert.New(t)

	// create empty map
	var map1 = map[string]int{}
	assert.Equal(0, len(map1))

	// set new elements in map
	map1["hour"] = 60 * 60 * 1000
	map1["minute"] = 60 * 1000
	map1["second"] = 1 * 1000
	assert.Equal(3, len(map1))

	// map returns result, isExist boolean
	var ms, ok = map1["minute"]
	if ok {
		assert.Equal(60*1000, ms)
	}

	delete(map1, "minute")
	_, ok = map1["minute"]
	assert.Equal(false, ok)

	var map2 = make(map[string]map[int]Student)
	map2["Software Engineers"] = map[int]Student{
		100: Student{
			"Anar",
			30,
		},

		200: Student{
			"Rufet",
			21,
		},
	}

	map2["Doctors"] = map[int]Student{
		100111: Student{
			"Araz",
			27,
		},

		100222: Student{
			"Elvira",
			26,
		},
	}

	assert.Equal("Anar", map2["Software Engineers"][100].name)

}

func TestHashTable(t *testing.T) {
	var res, err = http.Get("http://textfiles.com/groups/CHINA/china1.txt")
	if err != nil {
		log.Fatal(err)
	}

	var scanner = bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)
	var dictionary = make(map[string][]string)

	for scanner.Scan() {
		var text = scanner.Text()
		var firstLetter = HashBucket(text)

		if list, ok := dictionary[firstLetter]; ok {
			dictionary[firstLetter] = append(list, text)
		} else {
			list = []string{}
			dictionary[firstLetter] = append(list, text)
		}
	}

	pwd, _ := os.Getwd()
	var f, err1 = os.Create(pwd + "/dat1")
	check(err1)
	defer f.Close()
	w := bufio.NewWriter(f)

	for key, val := range dictionary {
		_, err2 := w.WriteString(key + ": [" + strings.Join(val, " ") + "]\n\n\n\n")
		check(err2)
	}

	w.Flush()
}

func HashBucket(word string) string {
	return strings.ToLower(string(word[0]))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (s Student) SetName(name string) {
	s.name = name
}

func (s *Student) SetNameWithPointer(name string) {
	s.name = name
}

func TestStructChangeName(t *testing.T) {
	var assert = assert.New(t)

	var s1 = Student{"Anar", 29}
	assert.Equal("Anar", s1.name)

	s1.SetName("Koko")
	assert.Equal("Anar", s1.name)

	s1.SetNameWithPointer("Momo")
	assert.Equal("Momo", s1.name)
}

type UserA struct {
	name string
}

type UserB struct {
	name string
}

type StudentA struct {
	UserA
	studentID int
}

func TestEmbeddedTypes(t *testing.T) {
	var assert = assert.New(t)

	var s = StudentA{
		UserA: UserA{
			name: "Anar",
		},
		studentID: 100,
	}

	assert.Equal("Anar", s.name)
}
