package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	var cmd string
	var err error
	reader := bufio.NewReader(os.Stdin)
	for {
		_, _ = os.Stdout.WriteString("$ ")
		cmd, err = reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading command:", err)
			return
		}
		cmd = cmd[:len(cmd)-1] // Remove the newline character
		fmt.Printf("%s: command not found\n", cmd)
	}
}
