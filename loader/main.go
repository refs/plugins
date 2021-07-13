package main

import (
	"log"
	"os"
	"os/signal"
	"plugin"
)

var plugins = []string{"../plugins/englishGreeter/main.so", "../plugins/chineseGreeter/main.so"}

func main() {
	// 1. load a plugin
	// 2. run the Serve function
	// 3. iterate over an array of plugins and call the serve function?
	// How can one enforce a contract on a plugin? A: Lookup should serve as enforcing a contract
	halt := make(chan os.Signal, 1)
	signal.Notify(halt, os.Interrupt)

	for i := range plugins {
		p, err := plugin.Open(plugins[i])
		if err != nil {
			log.Fatal(err)
		}

		s, err := p.Lookup("Serve")
		if err != nil {
			log.Fatal(err)
		}

		go s.(func())()
	}

	<-halt
}
