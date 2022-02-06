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

	ps "github.com/mitchellh/go-ps"
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
	if len(stderr.String()) != 0 {
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
	// log.Printf("Container IDs: %s\n", containerIds)

	var containers []Container

	var cmd, out string
	sep := ';'
	goTemplate := `{{ index .info.config.labels "io.kubernetes.pod.namespace"}}%c{{ index .info.config.labels "io.kubernetes.pod.name"}}%c{{ index .info.config.labels "io.kubernetes.container.name"}}%c{{.info.pid}}`
	goTemplate = fmt.Sprintf(goTemplate, sep, sep, sep)
	cmdTpl := `crictl inspect --output go-template --template '%v' %v`
	for _, cid := range containerIds {
		cmd = fmt.Sprintf(cmdTpl, goTemplate, cid)
		out, err = Shellout(cmd)
		if err != nil {
			return nil, err
		}

		containerDetails := strings.Split(out, string(sep))
		ns := containerDetails[0]
		pod := containerDetails[1]
		name := containerDetails[2]
		tmp := strings.TrimSpace(containerDetails[3])
		primaryPID, err := strconv.Atoi(tmp)
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

	// log.Println(containers)
	return containers, nil
}

func GetContainerFromPrimaryPid(containers []Container, pid int) (Container, bool) {
	for i := range containers {
		if containers[i].PrimaryPID == pid {
			return containers[i], true
		}
	}
	return Container{}, false
}

func GetPPid(pid int) int {
	log.Printf("PID %v", pid)
	p, err := ps.FindProcess(pid)
	if err != nil {
		panic(err)
	}
	return p.PPid()
}

func GetContainerFromPid(containers []Container, pid int) (Container, bool) {
	if pid == 1 {
		return Container{}, false
	}
	container, found := GetContainerFromPrimaryPid(containers, pid)
	if found {
		return container, true
	} else {
		pid = GetPPid(pid)
		return GetContainerFromPrimaryPid(containers, pid)
	}
}

func SplitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
