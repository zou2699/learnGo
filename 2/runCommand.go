package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmdOutput, e := exec.Command("go", "env").Output()
	if e!=nil {
		log.Fatal(e)
	}
	fmt.Printf("%s",cmdOutput)

}
