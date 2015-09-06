package jsongo

import (
	"testing"
	"encoding/json"
	"reflect"
	"log"
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

func bytes2json(data []byte) map[string]interface{} {
	var jsonData interface{}

	err := json.Unmarshal(data, &jsonData)

	if err != nil {
		log.Printf("Erro: %s", err)
		return nil
	}

	return jsonData.(map[string]interface{})
}

func struct2json(data interface{}) string {
	byteArray, err := json.Marshal(data)

	if err != nil {
		log.Printf("Erro: %s", err)
		return ""
	}

	return string(byteArray)
}

func check(t *testing.T, expect, result interface{}) {
	if !reflect.DeepEqual(expect, result) {
		t.Errorf("\n\nExpect: %s\nResult: %s", expect, result)
	}
}