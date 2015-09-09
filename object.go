package jsongo

import (
	"errors"
	"reflect"
	"fmt"
)

type O map[string]interface{}

func Object() O {
	return O{}
}

func (this O) Put(key string, value interface{}) O {
	this[key] = value
	return this
}

func (this O) Get(key string) interface{} {
	return this[key]
}

func (this O) GetObject(key string) (value O, err error) {
	switch this[key].(type) {
	case map[string]interface{}:
		object := Object()

		for k,v := range this[key].(map[string]interface{}) {
			object.Put(k,v)
		}

		return object, nil
	case O:
		return this[key].(O), nil
	}

	return nil, errors.New(fmt.Sprintf("Casting error. Interface is %s, not jsongo.object", reflect.TypeOf(this[key])))
}

func (this O) GetArray(key string) (newArray *A, err error) {
	newArray = Array()

	switch this[key].(type) {
	case []interface{}:
		values := this[key].([]interface{})

		for _, value := range values {
			newArray.Put(value)
		}

		return newArray, nil
	case []string:
		values := this[key].([]string)

		for _, value := range values {
			newArray.Put(value)
		}

		return newArray, nil
	case *A:
		return this[key].(*A), nil
	}

	return nil, errors.New(fmt.Sprintf("Casting error. Interface is %s, not jsongo.A or []interface{}", reflect.TypeOf(this[key])))
}

func (this O) Remove(key string) O {
	delete(this, key)
	return this
}

func (this O) Indent() string {
	return indent(this)
}

func (this O) String() string {
	return _string(this)
}