package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/atomgenie/konector/initialise"
	"github.com/atomgenie/konector/service"
	"github.com/atomgenie/konector/systemctl"
)

func main() {
	flag.Parse()
	argv := flag.Args()

	if len(argv) == 0 {
		fmt.Println("init | service | init-systemctl")
		flag.Usage()
		return
	}

	var err error

	switch argv[0] {
	case "init":
		err = initialise.Init(argv)
	case "service":
		err = service.StartService()
	case "init-systemctl":
		err = systemctl.Init(argv)
	}

	if err != nil {
		log.Fatalln(err)
	}

}
