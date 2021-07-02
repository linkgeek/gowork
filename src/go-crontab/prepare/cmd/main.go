package main

import (
	"fmt"
	"os/exec"
)

func demo1() {
	var (
		cmd *exec.Cmd
		err error
	)

	// cmd = exec.Command("/bin/bash", "-c", "echo 1;echo2;")
	// cmd = exec.Command("C:\\cygwin64\\bin\\bash.exe", "-c", "echo 1")
	cmd = exec.Command("F:\\Program Files\\Git\\git-bash.exe", "-c", "echo 1")

	err = cmd.Run()

	fmt.Println(err)
}

func main() {
	demo1()
}