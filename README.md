# sysdigtracers
Add sysdig tracers in your Go code easily

# sysdig tracers documentation 

https://github.com/draios/sysdig/wiki/Tracers

# usage

* Fetch and install package :
```
go get github.com/Issif/sysdigtracers
```
* Import package, add in your code :
```
import github.com/Issif/sysdigtracers
```
* Idiomatic way is to add at beginning of each function you want to trace :
```
t:= sysdigtracers.Entry("id", "tags", "args")
defer t.Exit("args")
```
* With (see https://github.com/draios/sysdig/wiki/Tracers#fields-explanation) :
    * **id** (required) : a string, can be empty
    * **tags** (required) : a string, if empty, will be set with name of caller function
    * **args** (optionnal) : a string, can be empty

* You can also add entry and exit events anywhere in your function :
```
func myFunction() {
//... some stuff ...
t := sysdigtracers.Entry("", "")
//... some stuff ...
t.Exit()
//... some stuff ...
}
```
# example

Inspired by : https://rosettacode.org/wiki/Mandelbrot_set#Go

```
package main

import (
	"fmt"
	"math/cmplx"

	"github.com/Issif/sysdigtracers"
)

func mandelbrot(a complex128) (z complex128) {
	t := sysdigtracers.Entry("", "")
	defer t.Exit("")
	for i := 0; i < 50; i++ {
		z = z*z + a
	}
	return
}

func main() {
	for y := 1.0; y >= -1.0; y -= 0.05 {
		for x := -2.0; x <= 0.5; x += 0.0315 {
			if cmplx.Abs(mandelbrot(complex(x, y))) < 2 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
```

```
~# sysdig evt.type=tracer

330655 11:38:17.059175478 1 mandelbrot (3034) > tracer id=0 tags=main.mandelbrot args=
330681 11:38:17.059254204 1 mandelbrot (3034) < tracer id=0 tags=main.mandelbrot args=
330718 11:38:17.059373437 1 mandelbrot (3034) > tracer id=0 tags=main.mandelbrot args=
330737 11:38:17.059472125 1 mandelbrot (3034) < tracer id=0 tags=main.mandelbrot args=
330772 11:38:17.059605109 1 mandelbrot (3034) > tracer id=0 tags=main.mandelbrot args=
330802 11:38:17.059685106 1 mandelbrot (3034) < tracer id=0 tags=main.mandelbrot args=
330839 11:38:17.059828369 1 mandelbrot (3034) > tracer id=0 tags=main.mandelbrot args=
```