package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func main() {

	ipAddress, err := exec.Command("curl", "ipinfo.io/ip").Output()
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	cred, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		os.Exit(0)
	}

	fmt.Printf("%s\n", cred)

	//config, err := google.ConfigFromJSON(cred, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	fmt.Printf("IP Address: %s\n", ipAddress)
}
