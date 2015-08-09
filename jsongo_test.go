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
	expect := bytes2json([]byte(`{"name":"Ricardo Longa","idade":28,"skills":["Golang","Android"]}`))
	result := Object().Put("name", "Ricardo Longa").Put("idade", 28).Put("skills", Array().Put("Golang").Put("Android"))

	check(t, struct2json(expect), struct2json(result))
}

func bytes2json(data []byte) map[string]interface{} {
	var jsonData interface{}

	err := json.Unmarshal(data, &jsonData)

	if err != nil {
		return nil
	}

	return jsonData.(map[string]interface{})
}

func struct2json(data interface{}) string {
	byteArray, err := json.Marshal(data)
	//	byteArray, err = json.MarshalIndent(pessoa, "", "    ")

	if err != nil {
		return ""
	}

	return string(byteArray)
}

func check(t *testing.T, expect, result interface{}) {
	if !reflect.DeepEqual(expect, result) {
		t.Errorf("\n\nExpect: %s\nResult: %s", expect, result)
	}

	log.Printf("")
}