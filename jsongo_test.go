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
	check(t, expect, result)
}

func Test_create_populated_object(t *testing.T) {
	expect := bytes2json([]byte(`{"name":"Ricardo Longa","idade":28,"skills":["Golang","Android"]}`))
	result := Object().Put("name", "Ricardo Longa").Put("idade", 28).Put("skills", Array().Put("Golang").Put("Android"))
	check(t, expect, result)
}

func Test_duas_strings(t *testing.T) {
	expect := "{\"name\":\"Ricardo Longa\",\"idade\":28,\"skills\":[\"Golang\",\"Android\"]}"
	result := "{\"skills\":[\"Android\",\"Golang\"]},\"idade\":28,\"name\":\"Ricardo Longa\"}"
	check(t, expect, result)
}

func bytes2json(data []byte) (map[string]interface{}) {
	var jsonData interface{}

	err := json.Unmarshal(data, &jsonData)

	if err != nil {
		return nil
	}

	return jsonData.(map[string]interface{})
}

func check(t *testing.T, expect, result interface{}) {
	log.Printf("Expect: %s\n", expect)
	log.Printf("Result: %s\n", result)

	if !reflect.DeepEqual(expect, result) {
		t.Errorf("\n\nExpect: %s\nResult: %s", expect, result)
	}

	log.Printf("")
}