package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const ShellToUse = "bash"

func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

type Container struct {
	Namespace  string
	Pod        string
	Id         string
	Name       string
	PrimaryPID int
}

func GetContainerDetails() error {
	var errOut error
	cmd := "crictl ps -q"
	err, out, stderr := Shellout(cmd)
	if err != nil {
		errMsg := fmt.Sprintf("error running command %v: %v", cmd, err)
		errOut = errors.New(errMsg)
		return errOut
	}
	if stderr != "" {
		errMsg := fmt.Sprintf("error returned by command %v: %v", cmd, stderr)
		errOut = errors.New(errMsg)
		return errOut
	}
	containerIds := SplitLines(out)

	for _, cid := range containerIds {
		goTemplate := `NS:{{ index .info.config.labels "io.kubernetes.pod.namespace"}} POD:{{ index .info.config.labels "io.kubernetes.pod.name"}} CONTAINER:{{ index .info.config.labels "io.kubernetes.container.name"}} PRIMARY PID:{{.info.pid}}`

		cmdTpl := `crictl inspect --output go-template --template '%v' %v`

		cmd = fmt.Sprintf(cmdTpl, goTemplate, cid)
		log.Println(cmd)

		err, out, stderr = Shellout(cmd)
		if err != nil {
			errMsg := fmt.Sprintf("error running command %v: %v", cmd, err)
			errOut = errors.New(errMsg)
		}
		log.Println(out)
		log.Println(stderr)
		log.Println(errOut)
	}

	return nil
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
