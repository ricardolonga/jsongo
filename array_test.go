package jsongo

import (
	"testing"
	"bytes"
	"strings"
)

func Test_create_two_populated_objects_into_array(t *testing.T) {
	expect := bytes2json([]byte(`{"funcionarios":[{"name":"Ricardo Longa","idade":28,"skills":["Golang","Android"]},{"name":"Hery Victor","idade":32,"skills":["Golang","Java"]}]}`))
	result := Object().Put("funcionarios", Array().Put(Object().Put("name", "Ricardo Longa").Put("idade", 28).Put("skills", Array().Put("Golang").Put("Android"))).
												   Put(Object().Put("name", "Hery Victor").Put("idade", 32).Put("skills", Array().Put("Golang").Put("Java"))))

	check(t, struct2json(expect), struct2json(result))
}

func Test_array_size_must_be_3(t *testing.T) {
	result := Array().Put("Golang").Put("Android").Put("Java")

	check(t, 3, result.Size())
}

func Test_range_array(t *testing.T) {
	results := Array().Put("Golang").Put("Android").Put("Java")

	expect := make(map[int]string, 0)
	expect[0] = "Golang"
	expect[1] = "Android"
	expect[2] = "Java"

	arrayOfString, _ := results.OfString()

	for i, result := range arrayOfString {
		if !strings.EqualFold(result, expect[i]) {
			t.Errorf("\n\nExpect: %s\nResult: %s", expect[i], result)
		}
	}
}

func Test_array_indent(t *testing.T) {
	expect := []byte(`[
   "Golang",
   "Android",
   "Java"
]`)
	result := Array().Put("Golang").Put("Android").Put("Java")

	if !bytes.Equal(expect, bytes.NewBufferString(result.Indent()).Bytes()) {
		t.Errorf("\n\nExpect: %s\nResult: %s", expect, result)
	}
}

func Test_array_string(t *testing.T) {
	expect := []byte(`["Golang","Android","Java"]`)
	result := Array().Put("Golang").Put("Android").Put("Java")

	if !bytes.Equal(expect, bytes.NewBufferString(result.String()).Bytes()) {
		t.Errorf("\n\nExpect: %s\nResult: %s", expect, struct2json(result.String()))
	}
}

func Test_error_expected_because_contains_a_int_value(t *testing.T) {
	array := Array().Put("Golang").Put(123).Put("Java")

	if _, err := array.OfString(); err == nil {
		t.Errorf("Error expected because this array contains a not string value.")
	}
}