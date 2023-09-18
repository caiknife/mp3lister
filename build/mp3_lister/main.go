package main

import (
	"log"
	"os"
)

func main() {
	if err := newApp().Run(os.Args); err != nil {
		log.Fatalln(err)
		return
	}
}
