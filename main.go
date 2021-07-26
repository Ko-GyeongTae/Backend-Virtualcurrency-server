package main

import (
	"github.com/Ko-GyeongTae/Backend-Virtualcurrency-server/cli"
	"github.com/Ko-GyeongTae/Backend-Virtualcurrency-server/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
