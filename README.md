# go-dms
A small library for converting Decimal Degrees to Degrees, Minutes, Seconds coordinates

Efficiently converting coordinates between DD and DMS

## Installation

`go get -u github.com/KingAkeem/go-dms/dms`

**test.go:**
```
package main

import (
    "github.com/KingAkeem/go-dms/dms"
    "fmt"
    "time"
    "log"
)

func main() {
    start := time.Now()
    dmsCoordinate, err := dms.New(2.21893, 1.213905)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("DMS coordinates:\n%+v, %+v\n", dmsCoordinate.Latitude.String(), dmsCoordinate.Longitude.String()) 
    end := time.Now()
    fmt.Printf("Function took %f seconds.\n", end.Sub(start).Seconds())
}
```
**>> go run test.go**

**Output:**
```
    DMS coordinates:
    2°13'8.148000" N, 1°12'50.058000" E
    Function took 0.000049 seconds.
```

**>> GOOS=js GOARCH=wasm go run -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" .** (Compiling as WebAssembly module for Node, run command in the same directory as `test.go`)

Golang WebAssemlby Wiki: https://github.com/golang/go/wiki/WebAssembly

**Output:**
```
    DMS coordinates:
    2°13'8.148000" N, 1°12'50.058000" E
    Function took 0.000478 seconds.
```


### [dms-js](https://github.com/WSDOT-GIS/dms-js)

**index.js:**
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

**>> node index.js**

**Output:** 
```
    DMS Coordinates: 2°13′8.147999999999422″ N, 1°12′50.058″ E
    Function took 0.001828968 seconds.
```


### [formatcoords](https://github.com/nerik/formatcoords)

**index.js:**
```
var formatcoords = require('formatcoords');
var NanoTimer = require('nanotimer');
var timerObj = new NanoTimer();

var dmsTest = function() {
    var coords = formatcoords(2.21893, 1.213905);
    console.log('DMS Coordinates: ' + coords.format());
    return coords;
};

var seconds = timerObj.time(dmsTest, "", 's')
console.log('Function took ' + seconds + ' seconds.');
```

**>> node index.js**

**Output:** 
```
DMS Coordinates: 2° 13′ 8.14800″ N 1° 12′ 50.05800″ E
Function took 0.002745904 seconds.
```





