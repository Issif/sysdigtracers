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
defer sysdigtracers.Entry("id", "tags", "args").Exit("args")
```
* With (see https://github.com/draios/sysdig/wiki/Tracers#fields-explanation) :
    * **id** (required) : a string, can be empty
    * **tags** (required) : a string, if empty, will be set with name of caller function
    * **args** (optionnal) : a string, can be empty

* You can also add entry and exit events anywhere in your function :
```
func myFunction() {
... some stuff ...
t := sysdigtracers.Entry("", "")
... some stuff ...
t.Exit()
... some stuff ...
}
```
# example

Inspired by : https://rosettacode.org/wiki/Mandelbrot_set#Go

```

```