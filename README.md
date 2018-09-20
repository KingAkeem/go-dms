# go-dms
A small library for converting Decimal Degrees to Degrees, Minutes, Seconds coordinates

This library provides efficient translation compared to others and will improved upon. I've currently tested against a JavaScript library.

test.go:
```
package main

import (
    "./dms"
    "fmt"
    "time"
    "log"
)

func main() {
    start := time.Now()
    lat, lon, err := dms.NewDMS(2.21893, 1.213905)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("DMS coordinates:\n%+v, %+v\n", lat.String(), lon.String()) 
    end := time.Now()
    fmt.Printf("Function took %f seconds.\n", end.Sub(start).Seconds())
}
```
**go run test.go**
**Output:**
```
    DMS coordinates:
    2°13'8.148000" N, 1°12'50.058000" E
    Function took 0.000049 seconds.
```

**>> GOOS=js GOARCH=wasm go run -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" .** (Compiling as WebAssembly module for Node, run command in the same directory as `test.go`)
**Output:**
```
    DMS coordinates:
    2°13'8.148000" N, 1°12'50.058000" E
    Function took 0.000478 seconds.
```


### [dms-js](https://github.com/WSDOT-GIS/dms-js)

index.js:
```
var dms = require('dms-conversion');
var NanoTimer = require('nanotimer');
var timerObj = new NanoTimer();

var dmsTest = function() {
    var coords = new dms.default(2.21893, 1.213905);
    console.log('DMS Coordinates: ' + coords.toString());
    return coords;
};

var seconds = timerObj.time(dmsTest, "", 's')
console.log('Function took ' + seconds + ' seconds.');
```

**node index.js**
**Output:** 
```
    DMS Coordinates: 2°13′8.147999999999422″ N, 1°12′50.058″ E
    Function took 0.001828968 seconds.
```




