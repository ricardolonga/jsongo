jsongo
===================

**Fluent API** to make it easier **to create Json** objects.

[![Build Status](https://travis-ci.org/ricardolonga/jsongo.svg)](https://travis-ci.org/ricardolonga/jsongo)[![](http://gocover.io/_badge/github.com/ricardolonga/jsongo)](http://gocover.io/github.com/ricardolonga/jsongo)

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
json := Object().Put("name", "Ricardo Longa").
				 Put("idade", 28).
				 Put("owner", true).
				 Put("skills", Array().Put("Golang").
									   Put("Android"))
```

License
-------------
Apache License Version 2.0, January 2004
