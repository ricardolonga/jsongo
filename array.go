package jsongo

import (
	"reflect"
	"errors"
	"fmt"
)

type array []interface{}

func Array() *array {
	return &array{}
}

func (this *array) Put(value interface{}) *array {
	*this = append(*this, value)
	return this
}

func (this *array) Indent() string {
	return indent(this)
}

func (this *array) String() string {
	return _string(this)
}

func (this *array) Size() int {
	return len(*this)
}

func (this *array) OfString() (values []string, err error) {
	for _, value := range *this {
		if reflect.TypeOf(value).String() != "string" {
			return nil, errors.New(fmt.Sprintf("Value is %s, not a string.", reflect.TypeOf(value)))
		}

		values = append(values, value.(string))
	}

	return values, nil
}