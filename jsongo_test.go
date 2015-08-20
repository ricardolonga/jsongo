package jsongo

import (
	"testing"
	"encoding/json"
	"reflect"
	"log"
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

func Test_create_two_populated_objects_into_array(t *testing.T) {
	expect := bytes2json([]byte(`{"funcionarios":[{"name":"Ricardo Longa","idade":28,"skills":["Golang","Android"]},{"name":"Hery Victor","idade":32,"skills":["Golang","Java"]}]}`))
	result := Object().Put("funcionarios", Array().Put(Object().Put("name", "Ricardo Longa").Put("idade", 28).Put("skills", Array().Put("Golang").Put("Android"))).
												   Put(Object().Put("name", "Hery Victor").Put("idade", 32).Put("skills", Array().Put("Golang").Put("Java"))))

	check(t, struct2json(expect), struct2json(result))
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