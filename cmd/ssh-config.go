package main

import (
	"fmt"
	"os"

	"github.com/kevinburke/ssh_config"
)

func main() {
	appName := "ssh-config"

	if len(os.Args) != 3 {
		if len(os.Args) >= 1 {
			appName = os.Args[0]
		}
		fmt.Printf("usage: %s <alias> <key>\n", appName)
		os.Exit(129)
	}

	println(ssh_config.Get(os.Args[1], os.Args[2]))
}