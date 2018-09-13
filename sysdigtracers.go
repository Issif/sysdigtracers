package sysdigtracers

import (
	"io/ioutil"
	"runtime"
)

// Tracer represents a tracer
type Tracer struct {
	// (see: https://github.com/draios/sysdig/wiki/Tracers#fields-explanation)

	// this string should be a 64bit integer or t, p, pp
	id string

	// this string should be a list of one or more strings separated by periods
	// if empty, name of caller function is set
	tags string

	// optional list of key-value pairs to be associated with the tracer
	args string
}

// getFunctionName retrieves name of caller function
func getFunctionName() string {
	fpcs := make([]uintptr, 1)
	runtime.Callers(3, fpcs)
	caller := runtime.FuncForPC(fpcs[0])
	return caller.Name()
}

// Entry emits an entry event in /dev/null and
// returns a tracer struct
// id, tags and args arguments can be empty string
// if tags is empty, name of caller function is set
func Entry(id string, element ...string) Tracer {
	if element[0] == "" {
		element[0] = getFunctionName()
	}

	var args string
	if len(element) > 1 {
		args = element[1]
	}

	t := Tracer{id: id, tags: element[0], args: args}

	d := []byte(">:" + t.id + ":" + t.tags + ":" + t.args + ":")
	ioutil.WriteFile("/dev/null", d, 0777)
	return t
}

// Exit emits an exit event in /dev/null
// args can be empty, if not, args of tracer are replaced
func (t Tracer) Exit(args ...string) {
	if len(args) != 0 {
		t.args = args[0]
	}

	d := []byte("<:" + t.id + ":" + t.tags + ":" + t.args + ":")
	ioutil.WriteFile("/dev/null", d, 0777)
}
