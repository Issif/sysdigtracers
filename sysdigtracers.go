// Package sysdigtracers let you add easily sysdig tracers in your Go code.
package sysdigtracers

import (
	"io/ioutil"
	"runtime"
	"strconv"
)

// Tracer represents a tracer
type Tracer struct {
	Id   string // Id string should be a 64bit integer or t, p, pp or empty (see: https://github.com/draios/sysdig/wiki/Tracers#fields-explanation).
	Tags string // Tags string should be a list of one or more strings separated by periods. if empty, name of caller function is set.
	Args string // Args string is a list of key-value pairs to be associated with the tracer (optionnal).
}

// getFunctionName retrieves name of caller function.
func getFunctionName() string {
	fpcs := make([]uintptr, 1)
	runtime.Callers(3, fpcs)
	caller := runtime.FuncForPC(fpcs[0])
	return caller.Name()
}

// Entry emits an entry event in /dev/null and returns a tracer struct.
// id and elements can be empty strings.
// first string in elements will be used as tags, other strings will be arguments.
// if elements is empty (no tag), name of caller function is set.
func Entry(id, tags string, args ...string) Tracer {
	switch id {
	case "", "t", "p", "pp":
		//
	default:
		if id != "" && id != "t" && id != "p" && id != "pp" {
			if _, err := strconv.Atoi(id); err != nil {
				id = ""
			}
		}
	}

	if tags == "" {
		tags = getFunctionName()
	}

	t := Tracer{Id: id, Tags: tags}

	if len(args) != 0 {
		t.Args = args[0]
	} else {
		t.Args = ""
	}

	d := []byte(">:" + t.Id + ":" + t.Tags + ":" + t.Args + ":")
	ioutil.WriteFile("/dev/null", d, 0777)
	return t
}

// Exit emits an exit event in /dev/null.
// args can be empty, if not, args of tracer are overrided by it.
func (t Tracer) Exit(args ...string) {
	if len(args) != 0 {
		t.Args = args[0]
	}

	d := []byte("<:" + t.Id + ":" + t.Tags + ":" + t.Args + ":")
	ioutil.WriteFile("/dev/null", d, 0777)
}
