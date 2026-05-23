package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	var cmd string
	var err error
	reader := bufio.NewScanner(os.Stdin)
	for {
		_, _ = os.Stdout.WriteString("$ ")
		reader.Scan()
		cmd = reader.Text()
		err = reader.Err()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading command:", err)
			return
		}
		if(cmd == "exit") {
			break
		} else if(strings.HasPrefix(cmd, "echo ")) {
			fmt.Println(cmd[len("echo "):])
		} else if(strings.HasPrefix(cmd, "type ")) {
			var builtIns = []string{"echo", "type", "exit"}
			if(!strings.Contains(strings.Join(builtIns, " "), cmd[len("type "):])) {
				fmt.Printf("%s: command not found\n", cmd[len("type "):])
			} else {
				fmt.Printf("%s is a shell builtin\n", cmd[len("type "):])
			}
		} else {
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
