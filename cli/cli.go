package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/Ko-GyeongTae/Backend-Virtualcurrency-server/explorer"
	"github.com/Ko-GyeongTae/Backend-Virtualcurrency-server/rest"
)

func usage() {
	fmt.Printf("Welcome to Neon Coin\n\n")
	fmt.Printf("Please se the following flags:\n\n")
	fmt.Printf("-port:			Set the PORT of the server\n")
	fmt.Printf("-mode: 			Choose between 'html' and 'rest'\n\n")
	fmt.Printf("-mode all -port:	Run REST API server on -port and Run HTML server on -port + 10\n\n")
	os.Exit(0)
}

func Start() {

	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of ther server")
	mode := flag.String("mode", "rest", "Choose in 'html', 'rest', 'all'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "all":
		go rest.Start(*port)
		explorer.Start(*port + 10)
	default:
		usage()
	}
}
