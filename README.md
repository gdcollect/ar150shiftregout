# ar150shiftregout
A Go package for use with the shift register 74HC595N for extending the number of outputs of the GL.iNet AR-150 router running OpenWrt.

## Installation
Install the ar150gpio package if you don't have it yet:
```
go get github.com/iketsj/ar150gpio
```

Then install it:
```
go get github.com/iketsj/ar150shiftregout
```

## Example Use
Note that the Initialize method have the parameters of dataInPinNumber, latchPinNumber, clockPinNumber, numOfShiftRegs respectively. If the numOfShiftRegs is equal to 2, the number of output pins would be 16.

```
package main

import (
	sr "github.com/iketsj/ar150shiftregout"
	"time"
)

func main() {
	myShiftRegister := sr.NewShiftRegister()
	myShiftRegister.Initialize(1, 14, 16, 17, 2)
	for i := uint8(0); i < 16; i++ {
		myShiftRegister.Write(i, sr.HIGH)
		myShiftRegister.Latch()
		time.Sleep(250 * time.Millisecond)
	}
}
```

## Building the source code
Do not forget to set the GOOS and GOARCH environment variable:
```
GOOS=linux GOARCH=mips go build main.go
```

## Transferring the Binary
You could by using scp:
```
scp main root@192.168.8.1:/
```

## Links
* http://www.ti.com/lit/ds/symlink/sn74hc595.pdf
* https://www.gl-inet.com/products/gl-ar150/


