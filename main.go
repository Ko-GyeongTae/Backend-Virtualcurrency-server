package main

import (
	"github.com/Ko-GyeongTae/Backend-Virtualcurrency-server/explorer"
	"github.com/Ko-GyeongTae/Backend-Virtualcurrency-server/rest"
)

func main() {
	explorer.Start(3000)
	rest.Start(4000)
}
