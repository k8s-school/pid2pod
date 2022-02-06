package main

import (
	"flag"
)

func main() {

	// pid := flag.Int("p", 0, "PID number on host")

	flag.Parse()

	GetContainerDetails()
}
