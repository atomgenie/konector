package main

import (
	"flag"
	"log"

	"github.com/atomgenie/konector/initialise"
	"github.com/atomgenie/konector/service"
)

func main() {
	flag.Parse()
	argv := flag.Args()

	if len(argv) == 0 {
		flag.Usage()
		return
	}

	var err error

	switch argv[0] {
	case "init":
		err = initialise.Init(argv)
	case "service":
		err = service.StartService()
	}

	if err != nil {
		log.Fatalln(err)
	}

}
