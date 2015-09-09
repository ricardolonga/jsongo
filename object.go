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
	var ok bool

	if value, ok = this[key].(O); !ok {
		return nil, errors.New(fmt.Sprintf("Casting error. Interface is %s, not jsongo.object", reflect.TypeOf(this[key])))
	}

	return value, nil
}

func (this O) GetArray(key string) (value *A, err error) {
	var ok bool

	if value, ok = this[key].(*A); !ok {
		return nil, errors.New(fmt.Sprintf("Casting error. Interface is %s, not jsongo.array", reflect.TypeOf(this[key])))
	}

	return value, nil
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