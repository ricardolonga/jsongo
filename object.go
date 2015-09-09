package jsongo

import (
	"errors"
	"reflect"
	"fmt"
)

type object map[string]interface{}

func Object() object {
	return object{}
}

func (this object) Put(key string, value interface{}) object {
	this[key] = value
	return this
}

func (this object) Get(key string) interface{} {
	return this[key]
}

func (this object) GetObject(key string) (value object, err error) {
	var ok bool

	if value, ok = this[key].(object); !ok {
		return nil, errors.New(fmt.Sprintf("Casting error. Interface is %s, not jsongo.object", reflect.TypeOf(this[key])))
	}

	return value, nil
}

func (this object) GetArray(key string) (value *array, err error) {
	var ok bool

	if value, ok = this[key].(*array); !ok {
		return nil, errors.New(fmt.Sprintf("Casting error. Interface is %s, not jsongo.array", reflect.TypeOf(this[key])))
	}

	return value, nil
}

func (this object) Remove(key string) object {
	delete(this, key)
	return this
}

func (this object) Indent() string {
	return indent(this)
}

func (this object) String() string {
	return _string(this)
}