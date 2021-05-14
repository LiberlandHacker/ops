package cmd

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/nanovms/ops/util/slice"
)

// Checks which process is using given port number
func checkPortUserPID(portNumber string) (string, error) {
	var (
		pid string
		cmd = exec.Command("lsof", "-i", ":"+portNumber)
	)
	out, err := cmd.Output()
	if err != nil {
		return pid, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	scanner.Scan() // check header first
	cols := slice.ExcludeWhitespaces(strings.Split(scanner.Text(), " "))
	headerIdx := -1
	for i, v := range cols {
		if strings.ToLower(v) == "pid" {
			headerIdx = i
			break
		}
	}

	scanner.Scan()
	line := strings.Trim(scanner.Text(), " ")
	if line != "" {
		cols = slice.ExcludeWhitespaces(strings.Split(line, " "))
		pid = cols[headerIdx]
	}

	return pid, err
}
