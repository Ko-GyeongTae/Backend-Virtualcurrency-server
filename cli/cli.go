package cli

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Welcome to Neon Coid\n\n")
	fmt.Printf("Please se the following Commands:\n\n")
	fmt.Printf("explorer:	Start the HTML Explorer\n")
	fmt.Printf("rest:		Start the REST API (recommended)\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) < 2 {
		usage()
	}

	rest := flag.NewFlagSet("rest", flag.ExitOnError)

	portFlag := rest.Int("port", 4000, "Sets ther port of the server")

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		fmt.Println("Start REST API")
		rest.Parse(os.Args[2:])
	default:
		usage()
	}

	if rest.Parsed() {
		fmt.Println(portFlag)
		fmt.Println("Start server")
	}
}
