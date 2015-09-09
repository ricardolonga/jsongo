package jsongo

import (
	"testing"
	"bytes"
	"strings"
)

func Test_create_empty_object(t *testing.T) {
	expect := bytes2json([]byte(`{}`))
	result := Object()

	check(t, struct2json(expect), struct2json(result))
}

func Test_create_populated_object(t *testing.T) {
	expect := bytes2json([]byte(`{"name":"Ricardo Longa","idade":28,"owner":true,"skills":["Golang","Android"]}`))
	result := Object().Put("name", "Ricardo Longa").Put("idade", 28).Put("owner", true).Put("skills", Array().Put("Golang").Put("Android"))

	check(t, struct2json(expect), struct2json(result))
}

func Test_create_populated_objects_and_remove_attr(t *testing.T) {
	expect := bytes2json([]byte(`{"name":"Ricardo Longa","idade":28,"skills":["Golang","Android"]}`))
	result := Object().Put("name", "Ricardo Longa").Put("idade", 28).Put("skills", Array().Put("Golang").Put("Android"))

	check(t, struct2json(expect), struct2json(result))

	expectAfterRemove := bytes2json([]byte(`{"name":"Ricardo Longa","idade":28}`))

	result.Remove("skills")

	check(t, struct2json(expectAfterRemove), struct2json(result))
}

func Test_object_get_func(t *testing.T) {
	expect := "Ricardo Longa"
	result := Object().Put("name", "Ricardo Longa")

	if !strings.EqualFold(expect, result.Get("name").(string)) {
		t.Errorf("\n\nExpect: %s\nResult: %s", expect, result.Get("name"))
	}
}

func Test_object_indent(t *testing.T) {
	expect := []byte(`{
   "skills": [
      "Golang",
      "Android",
      "Java"
   ]
}`)
	result := Object().Put("skills", Array().Put("Golang").Put("Android").Put("Java"))

	if !bytes.Equal(expect, bytes.NewBufferString(result.Indent()).Bytes()) {
		t.Errorf("\n\nExpect: %s\nResult: %s", expect, struct2json(result.Indent()))
	}
}

func Test_object_string(t *testing.T) {
	expect := []byte(`{"skills":["Golang","Android","Java"]}`)
	result := Object().Put("skills", Array().Put("Golang").Put("Android").Put("Java"))

	if !bytes.Equal(expect, bytes.NewBufferString(result.String()).Bytes()) {
		t.Errorf("\n\nExpect: %s\nResult: %s", expect, struct2json(result.String()))
	}
}

func Test_get_object_with_casting_error(t *testing.T) {
	obj := Object().Put("skills", Array().Put("Golang").Put("Android").Put("Java"))

	if _, err := obj.GetObject("skills"); err == nil {
		t.Errorf("Casting error not found.")
	}
}

func Test_get_object_without_casting_error(t *testing.T) {
	obj := Object().Put("owner", Object().Put("nome", "Ricardo Longa"))

	if _, err := obj.GetObject("owner"); err != nil {
		t.Errorf("Casting error not expected.")
	}
}

func Test_get_array_without_casting_error(t *testing.T) {
	obj := Object().Put("skills", Array().Put("Golang").Put("Android").Put("Java"))

	values, err := obj.GetArray("skills")
	if err != nil {
		t.Errorf("Error not expected: %s.", err)
	}

	if len(*values) != 3 {
		t.Error("Expected 3 values.")
	}

	obj = Object().Put("skills", []interface{}{"Golang", "Android", "Java"})

	values, err = obj.GetArray("skills")
	if err != nil {
		t.Errorf("Error not expected: %s.", err)
	}

	if len(*values) != 3 {
		t.Error("Expected 3 values.")
	}

	obj = Object().Put("skills", []string{"Golang", "Android", "Java"})

	values, err = obj.GetArray("skills")
	if err != nil {
		t.Errorf("Error not expected: %s.", err)
	}

	if len(*values) != 3 {
		t.Error("Expected 3 values.")
	}
}

func Test_get_array_with_casting_error(t *testing.T) {
	obj := Object().Put("owner", Object().Put("nome", "Ricardo Longa"))

	if _, err := obj.GetArray("owner"); err == nil {
		t.Errorf("Casting error not found.")
	}
}