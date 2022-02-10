
<p align="center">
<b style="font-size: 30px">JSON SPANNER</b>
<p>
JSON Spanner is a Go package that provides a fast and simple way to filter or transform a json document. You can filter or transform json using SQL-like statements.

Getting Started
===============

## Installing

To start using JSON Spanner, install Go and run `go get`:

```sh
$ go get -u github.com/fuhongbo/json_spanner
```

This will retrieve the library.

## Filter a json document

Use SQL-like statements to filter JSON, return true if conditions are met, and return false if not.

```go
package main

import (
	"fmt"
	"github.com/fuhongbo/json_spanner"
)

func main() {
	eng, _ := NewFilter("test.a.b=1 and (b=2 or c=4)")
	result := eng.Valid(`{"test":{"a":{"b":1}},"b":3,"c":4}`)
	fmt.Println(result)
}
```

This will print:

```
true
```

## Transform a json document

Use SQL-like statements to transform JSON.

```go
package main

import "fmt"

func main() {
	eng, err := NewQuery("select test.a as x.t, 6.8 as x.d,b as x.b,c where test.a=1 and (b=2 or c=4)")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	match, jsonStr, err := eng.Transform(`{"test":{"a":1},"b":3,"c":4}`)
    
	if !match{
		fmt.Println("the json document do not match the condition")
		return
    } 
	
	if err!=nil{
		fmt.Println(err.Error())
		return
    }
	fmt.Println(jsonStr)
}
```

This will print:

```json
{
    "c": 4,
    "x": {
        "b": 3,
        "d": 6.8,
        "t": 1
    }
}
```

## Currently supported conditional expressions 

| Expressions | Remark                               |
|-------------|--------------------------------------|
| =           |                                      |
| <>          |                                      |
| \>          |                                      |
| \>=         |                                      |
| <           |                                      |
| <=          |                                      |
| AND         |                                      |
| OR          |                                      |
| =~          | This special use to MQTT topic match |
