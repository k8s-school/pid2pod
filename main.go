package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "PID")
		return
	}
	pid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: incorrect argument: %v\n", err)
		os.Exit(1)
	}
	containers, err := GetContainerDetails()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	container, found, err := GetContainerFromPid(containers, pid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	} else if !found {
		fmt.Fprintf(os.Stderr, "error: process %d is running on host\n", pid)
		os.Exit(1)
	}
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "NAMESPACE\tPOD\tCONTAINER\tPRIMARY PID")
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", container.Namespace, container.Pod, container.Name, container.PrimaryPID)
	w.Flush()

}
