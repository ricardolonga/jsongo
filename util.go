package jsongo

import "encoding/json"

func indent(v interface{}) string {
	indent, _ := json.MarshalIndent(v, "", "   ")
	return string(indent)
}

func _string(v interface{}) string {
	indent, _ := json.Marshal(v)
	return string(indent)
}
