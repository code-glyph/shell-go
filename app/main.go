package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	_, _ = os.Stdout.WriteString("$ ")
	var cmd string
	var err error
	cmd, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading command:", err)
		return
	}
	cmd = cmd[:len(cmd)-1] // Remove the newline character
	fmt.Printf("%s: command not found\n", cmd)
}
