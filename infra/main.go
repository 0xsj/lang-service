package main

import (
	"fmt"
	"os"

	"github.com/0xsj/lang-service/infra/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(("Usage: go run main.go [deploy]"))
		return
	}

	switch os.Args[1] {
	case "deploy":
		cmd.Deploy()
	default:
		fmt.Println(("Invalid command. Use 'deploy"))
	}
}