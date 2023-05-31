package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ssh", "-C -D 12580", "vps")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
