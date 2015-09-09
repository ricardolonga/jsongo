package jsongo

import (
	"log"
	"encoding/json"
	"reflect"
	"testing"
)

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