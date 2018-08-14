package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	ipAddress, err := exec.Command("curl", "ipinfo.io/ip").Output()

	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	fmt.Printf("IP Address: %s\n", ipAddress)
}
