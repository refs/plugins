package main

import (
	"log"
	"os"
	"os/signal"
	"plugin"
)

var plugins = []string{"../plugins/englishGreeter/main.so", "../plugins/chineseGreeter/main.so"}

func main() {
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
