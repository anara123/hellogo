package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Car struct {
	Brand           string
	Model           string
	Number          string `json:"number_tag"`
	notExportedProp string `json:"notExportedProp"` // this prop is ignored
}

func TestMarshal(t *testing.T) {
	var assert = assert.New(t)

	var car = Car{"BMW", "V7", "90FB900", "foo"}
	var bytes, _ = json.Marshal(&car)

	assert.Equal(`{"Brand":"BMW","Model":"V7","number_tag":"90FB900"}`, string(bytes))
}

func TestUnmarshal(t *testing.T) {
	var assert = assert.New(t)

	var car Car
	var err = json.Unmarshal([]byte(`{"Brand":"BMW","Model":"V7","number_tag":"90FB900"}`), &car)

	assert.Nil(err)
	assert.Equal(car, Car{"BMW", "V7", "90FB900", ""})
}

func TestUnmarshalJsonHasMorePropertiesThanInStruct(t *testing.T) {
	var assert = assert.New(t)

	var car Car
	var err = json.Unmarshal([]byte(`{"Brand":"BMW","Model":"V7","number_tag":"90FB900","additional_prop":500}`), &car)

	assert.Nil(err)
	assert.Equal(car, Car{"BMW", "V7", "90FB900", ""})
}

func TestUnmarshalJsonHasLessPropsThanInStruct(t *testing.T) {
	var assert = assert.New(t)

	var car Car
	var err = json.Unmarshal([]byte(`{"Brand":"BMW","Model":"V7"}`), &car)

	assert.Nil(err)
	assert.Equal(Car{"BMW", "V7", "", ""}, car)
}

func ExampleEncodeJSON() {
	var car = Car{"BMW", "V7", "90FB900", "foo"}
	json.NewEncoder(os.Stdout).Encode(&car)

	// Output:
	// {"Brand":"BMW","Model":"V7","number_tag":"90FB900"}
}

func TestDecodeJSON(t *testing.T) {
	var assert = assert.New(t)

	var car Car
	var reader = strings.NewReader(`{"Brand":"BMW","Model":"V7","number_tag":"90FB900"}`)
	var err = json.NewDecoder(reader).Decode(&car)

	assert.Nil(err)
	assert.Equal(Car{"BMW", "V7", "90FB900", ""}, car)
}
