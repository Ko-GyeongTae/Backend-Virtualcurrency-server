package main

import (
	"github.com/Ko-GyeongTae/Backend-Virtualcurrency-server/explorer"
	"github.com/Ko-GyeongTae/Backend-Virtualcurrency-server/rest"
)

func main() {
	go explorer.Start(4000)
	rest.Start(4001)
}
