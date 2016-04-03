package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Product struct {
	number int
	name   string
}

type ByName []Product

func (b ByName) Len() int {
	return len(b)
}

func (b ByName) Swap(i, j int) {
	var temp = b[i]
	b[i] = b[j]
	b[j] = temp
}

func (b ByName) Less(i, j int) bool {
	return b[i].name < b[j].name
}

func (b ByName) Sort() []Product {
	sort.Sort(b)
	return b
}

func TestSortStructsByDifferentParams(t *testing.T) {
	var assert = assert.New(t)

	var arr = []Product{
		Product{111, "IPhone"},
		Product{222, "Android"},
		Product{333, "Sony"},
		Product{444, "Motorolla"},
	}

	sort.Sort(ByName(arr))

	assert.Equal(
		[]Product{
			Product{222, "Android"},
			Product{111, "IPhone"},
			Product{444, "Motorolla"},
			Product{333, "Sony"},
		},
		arr)

	assert.Equal(
		[]Product{
			Product{222, "Android"},
			Product{111, "IPhone"},
			Product{444, "Motorolla"},
			Product{333, "Sony"},
		},
		ByName(arr).Sort())
}

func TestSortStringSlice(t *testing.T) {
	var assert = assert.New(t)

	var arr = []string{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(sort.StringSlice(arr))

	assert.Equal([]string{"Al", "Jenny", "John", "Zeno"}, arr)
}

func TestReverseSortStringSlice(t *testing.T) {
	var assert = assert.New(t)

	var arr = []string{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(sort.Reverse(sort.StringSlice(arr)))

	assert.Equal([]string{"Zeno", "John", "Jenny", "Al"}, arr)
}

func TestSortNumberSlice(t *testing.T) {
	// var assert = assert.New(t)
	var assert = assert.New(t)

	var arr = []int{44, 12, 55, 123, 1, 23}
	sort.Sort(sort.IntSlice(arr))

	assert.Equal([]int{1, 12, 23, 44, 55, 123}, arr)
}
