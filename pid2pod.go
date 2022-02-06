package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const ShellToUse = "bash"

type Container struct {
	Namespace  string
	Pod        string
	Id         string
	Name       string
	PrimaryPID int
}

func Shellout(command string) (string, error) {
	var errOut error
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		errMsg := fmt.Sprintf("error running command %v: %v", cmd, err)
		errOut = errors.New(errMsg)
		return "", errOut
	}
	if len(stderr.String()) == 0 {
		errMsg := fmt.Sprintf("error returned by command %v: %v", cmd, stderr)
		errOut = errors.New(errMsg)
		return "", errOut
	}
	return stdout.String(), nil
}

func GetContainerIds() ([]string, error) {
	cmd := "crictl ps -q"
	out, err := Shellout(cmd)
	if err != nil {
		return nil, err
	}
	containerIds := SplitLines(out)
	return containerIds, nil
}

func GetContainerDetails() ([]Container, error) {

	containerIds, err := GetContainerIds()
	if err != nil {
		return nil, err
	}

	var containers []Container

	var cmd, out string
	sep := ';'
	goTemplate := `{{ index .info.config.labels "io.kubernetes.pod.namespace"}}%1$s{{ index .info.config.labels "io.kubernetes.pod.name"}}%1$s{{ index .info.config.labels "io.kubernetes.container.name"}}%1$s{{.info.pid}}`
	goTemplate = fmt.Sprintf(goTemplate, sep)
	cmdTpl := `crictl inspect --output go-template --template '%v' %v`
	for _, cid := range containerIds {

		cmd = fmt.Sprintf(cmdTpl, goTemplate, cid)
		out, err = Shellout(cmd)
		if err != nil {
			return nil, err
		}
		log.Println(out)

		containerDetails := strings.Split(out, string(sep))
		ns := containerDetails[0]
		pod := containerDetails[1]
		name := containerDetails[2]
		primaryPID, err := strconv.Atoi(containerDetails[3])
		if err != nil {
			return nil, err
		}
		container := Container{
			ns,
			pod,
			cid,
			name,
			primaryPID,
		}
		containers = append(containers, container)
	}
	return containers, nil
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
