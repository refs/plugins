package main

import (
	"log"
	"plugin"
)

func main() {
	// load a plugin
	// run the Start function
	// TODO: how can one enforce a contract on a plugin?

	p, err := plugin.Open("../plugins/server.so")
	if err != nil {
		log.Fatal(err)
	}

	s, err := p.Lookup("Serve")
	if err != nil {
		log.Fatal(err)
	}

	s.(func())()
}
