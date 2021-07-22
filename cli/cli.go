package cli

import (
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

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		fmt.Println("Start REST API")
	default:
		usage()
	}
}
