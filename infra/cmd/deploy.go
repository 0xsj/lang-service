package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func Deploy() {
	fmt.Println("Initializing terraform")
	cmd := exec.Command("terraform", "-chdir=infra/terraform", "init")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Terraform init failed", err)
		return
	}

	fmt.Println("applying terraform")
}
