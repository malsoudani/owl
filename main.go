package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmdArr := strings.Fields("ps")
	// clear := func() {
	// 	cmd := exec.Command("clear")
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Run()
	// }
	for {
		cmd := exec.Command(cmdArr[0], cmdArr[1:]...)
		if cmd.Stdout != os.Stdout {
			cmd.Stdout = os.Stdout
			if err := cmd.Run(); err != nil {
				fmt.Println(os.Stderr, err)
			}
		}
		cmd.Stdout = os.Stdout
		// clear()
	}
}
