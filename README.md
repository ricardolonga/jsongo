Jsongo
===================

**Fluent API** to make it easier **to create Json** objects.

[![travis-ci](https://travis-ci.org/ricardolonga/jsongo.svg)](https://travis-ci.org/ricardolonga/jsongo) 
[![codecov](https://codecov.io/gh/ricardolonga/jsongo/branch/master/graph/badge.svg)](https://codecov.io/gh/ricardolonga/jsongo)
[![goreportcard](https://goreportcard.com/badge/github.com/ricardolonga/jsongo)](http://gocover.io/github.com/ricardolonga/jsongo)

Install
-------------

```sh
go get github.com/ricardolonga/jsongo
```

Usage
-------------

To create this:  

```json
{  
    "name":"Ricardo Longa",
    "idade":28,
    "owner":true,
    "skills":[  
        "Golang",
        "Android"
    ]
}
```

Do this:  

```go
import (
    j "github.com/ricardolonga/jsongo"
)

json := j.Object().
		Put("name", "Ricardo Longa").
		Put("idade", 28).
		Put("owner", true).
		Put(
			"skills",
			j.Array().
				Put("Golang").
				Put("Android"),
		)

log.Println(json.Indent())
log.Println(json.String())
```

##### Convert object/array to indented String:

```go
json.Indent()
```

##### Convert object/array to String:

```go
json.String()
```

##### To remove a field of the object:

```go
json.Remove("skills")
```

##### To get a field of the object:

```go
json.Get("skills") // Return is interface{}.
```

##### To range over a array:

```go
results := Array().Put("Golang").Put("Android").Put("Java")

for i, result := range results.Array() {
}
```

##### To get Array size:

```go
array := j.Array().Put("Android").
                   Put("Golang").
                   Put("Java")
                   
array.Size() // Result is 3.
```

Copyright
-------------
Copyright (c) 2015 Ricardo Longa.  
Jsongo is licensed under the **Apache License Version 2.0**. See the LICENSE file for more information.
