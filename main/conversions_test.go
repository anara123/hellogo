package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToFloat(t *testing.T) {
	var assert = assert.New(t)

	var i = 5
	var f = 5.125

	assert.Equal(10.125, float64(i)+f)
}

func TestFloatToInt(t *testing.T) {
	var assert = assert.New(t)

	var i = 5
	var f = 5.125

	assert.Equal(10, i+int(f))
}

func TestRuneToString(t *testing.T) {
	var assert = assert.New(t)

	var r rune = 'a'
	var r2 int32 = 'b'

	assert.Equal("a", string(r))
	assert.Equal("b", string(r2))
}

func TestRuneToSliceOfBytesToString(t *testing.T) {
	var assert = assert.New(t)

	assert.Equal("hello", string([]byte{'h', 'e', 'l', 'l', 'o'}))
}

func TestStringToSliceOfBytes(t *testing.T) {
	assert.Equal(t,
		[]byte{0x68, 0x65, 0x6c, 0x6c, 0x6f},
		[]byte("hello"))
}

func TestAsciiToInt(t *testing.T) {

}
