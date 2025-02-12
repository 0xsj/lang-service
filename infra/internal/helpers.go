package internal

import (
	"fmt"
	"os/exec"
)

func RunCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = exec.Command("tee", "/dev/stderr").Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("Command failed: %s\n", err)
	}
	return err
}