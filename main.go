package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	cmdArr := strings.Fields("ps")
	clear := func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	for {
		cmd := exec.Command(cmdArr[0], cmdArr[1:]...)
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Println(os.Stderr, err)
		}
		cmd.Stdout = os.Stdout
		time.Sleep(3 * time.Second)
		clear()
	}
}
