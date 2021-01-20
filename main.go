package main

import (
	"math/rand"
	"time"

	"github.com/donech/erp/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cmd.Execute()
}
