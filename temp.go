package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
	piTemp()
}

func piTemp() {
	cmd, err := exec.Command("vcgencmd", "measure_temp").Output()
	if err != nil {
		fmt.Println(err)
	}
	temp := strings.ReplaceAll(string(cmd), "temp=", "")
	for {
		fmt.Printf(temp)
		time.Sleep(1000 * time.Millisecond)
	}
}
