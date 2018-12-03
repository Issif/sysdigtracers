package sysdigtracers_test

import "sysdigtracers"

// Idiomatic way to add a tracer.
func ExampleEntry() {
	t := sysdigtracers.Entry("id", "tags", "args")
	defer t.Exit("args")
}

// Add a tracer anywhere.
func ExampleEntry_other() {
	//... some stuff ...
	t := sysdigtracers.Entry("", "")
	//... some stuff ...
	t.Exit()
	//... some stuff ...
}

// Add a tracer in a sublevel goroutine.
func ExampleEntry_goroutine() {
	//... some stuff ...
	t := sysdigtracers.Entry("", "root")
	go func() {
		u := sysdigtracers.Entry("", "root.goroutine")
		defer u.Exit()
		//... some stuff ...
	}()
	t.Exit()
	//... some stuff ...
}
