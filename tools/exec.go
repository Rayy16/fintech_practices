package tools

import (
	"bytes"
	"os/exec"
	"strings"
)

func SplitCmd(cmdStr string) (name string, args []string) {
	if cmdStr == "" {
		return "", nil
	}
	cmds := strings.Split(cmdStr, " ")
	return cmds[0], cmds[1:]
}

func RunCmd(name string, args []string) (std [2]string, err error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	std = [2]string{stdout.String(), stderr.String()}
	return
}
