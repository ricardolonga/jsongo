package jsongo

import (
	"reflect"
	"errors"
	"fmt"
)

type A []interface{}

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
	return _string(this)
}

func (this *A) Size() int {
	return len(*this)
}

func (this *A) OfString() (values []string, err error) {
	for _, value := range *this {
		if reflect.TypeOf(value).String() != "string" {
			return nil, errors.New(fmt.Sprintf("Value is %s, not a string.", reflect.TypeOf(value)))
		}

		values = append(values, value.(string))
	}

	return values, nil
}