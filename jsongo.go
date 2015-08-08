package jsongo

type O map[string]interface{}

type A []interface{}

func Object() O {
	return O{}
}

func (this O) Put(key string, value interface{}) O {
	this[key] = value
	return this
}

func Array() A {
	return A{}
}

func (this A) Put(value interface{}) A {
	this = append(this, value)
	return this
}