package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	pid := flag.Int("p", 0, "PID number on host")

	flag.Parse()

	containers, err := GetContainerDetails()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	container, found := GetContainerFromPid(containers, *pid)
	if !found {
		fmt.Fprintf(os.Stderr, "process %v is running on host", pid)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%v", container)
}
