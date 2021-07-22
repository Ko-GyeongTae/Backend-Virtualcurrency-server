package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/Ko-GyeongTae/Backend-Virtualcurrency-server/explorer"
	"github.com/Ko-GyeongTae/Backend-Virtualcurrency-server/rest"
)

func usage() {
	fmt.Printf("Welcome to Neon Coid\n\n")
	fmt.Printf("Please se the following flags:\n\n")
	fmt.Printf("-port:		Set the PORT of the server\n")
	fmt.Printf("-mode: 		Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {

	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of ther server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
