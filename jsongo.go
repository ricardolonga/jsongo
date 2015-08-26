package jsongo

import "encoding/json"

type O map[string]interface{}

type A []interface{}

func Object() O {
	return O{}
}

func (this O) Put(key string, value interface{}) O {
	this[key] = value
	return this
}

func (this O) Remove(key string) O {
	delete(this, key)
	return this
}

func (this O) Indent() string {
	return indent(this)
}

func (this O) String() string {
	return toString(this)
}

func Array() *A {
	return &A{}
}

func (this *A) Put(value interface{}) *A {
	*this = append(*this, value)
	return this
}

func (this *A) Indent() string {
	return indent(this)
}

func (this *A) String() string {
	return toString(this)
}

func (this *A) Size() int {
	return len(*this)
}

func indent(v interface{}) string {
	indent, _ := json.MarshalIndent(v, "", "   ")
	return string(indent)
}

func toString(v interface{}) string {
	indent, _ := json.Marshal(v)
	return string(indent)
}
