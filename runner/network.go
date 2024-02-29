package runner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

func RunCommandAndDecode[T any](command string, args []string) (*T, error) {
	fullCommand := []string{"/usr/bin/vtysh", "-c", command}
	cmd := exec.Command("sudo", fullCommand...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("command err: %s", stderr.String())
	}

	var result T
	err = json.Unmarshal(out.Bytes(), &result)
	if err != nil {
		return nil, fmt.Errorf("JSON decode hatasi: %w", err)

	}
	return &result, nil
}
