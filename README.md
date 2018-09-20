# go-dms
A small library for converting Decimal Degrees to Degrees, Minutes, Seconds coordinates

This library provides efficient translation compared to others and will improved upon. I've currently tested against a JavaScript library by building `go-dms` as a WebAssembly module.

### [dms-js](https://github.com/WSDOT-GIS/dms-js)
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

```
Output:  
DMS Coordinates: 2°13′8.147999999999422″ N, 1°12′50.058″ E
Function took 0.001828968 seconds.
```


