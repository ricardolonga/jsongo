Jsongo
===================

**Fluent API** to make it easier **to create Json** objects.

[![Build Status](https://travis-ci.org/ricardolonga/jsongo.svg)](https://travis-ci.org/ricardolonga/jsongo) 
[![](http://gocover.io/_badge/github.com/ricardolonga/jsongo)](http://gocover.io/github.com/ricardolonga/jsongo)

Install
-------------
```
go get github.com/ricardolonga/jsongo
```

Usage
-------------
To create this:  
```
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
```
import (
    j "github.com/ricardolonga/jsongo"
)

json := j.Object().Put("name", "Ricardo Longa").
				   Put("idade", 28).
				   Put("owner", true).
				   Put("skills", j.Array().Put("Golang").
									       Put("Android"))

log.Print(json.Indent())
```

Copyright
-------------
Copyright (c) 2015 Ricardo Longa.  
Jsongo is licensed under the **Apache License Version 2.0**. See the LICENSE file for more information.
