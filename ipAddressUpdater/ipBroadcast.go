package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {

	ipAddress, err := exec.Command("curl", "ipinfo.io/ip").Output()
	//credentials := getCredentials

	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	fmt.Printf("IP Address: %s\n", ipAddress)
}

func getCredentials() []byte {
	cred, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		os.Exit(0)
	}
	return cred
}
