package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"

	"encoding/json"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	json.NewEncoder(f).Encode(token)
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func main() {
	ctx := context.Background()
	address, err := exec.Command("curl", "ipinfo.io/ip").Output()
	var ipAddress string
	ipAddress = string(address)
	fmt.Printf("Got IP Address to: %s", ipAddress)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	cred, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
		os.Exit(0)
	}

	config, err := google.ConfigFromJSON(cred, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)
	fmt.Printf("Begine Sheets\n")
	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	fmt.Printf("Set spreadSheet data\n")
	spreadsheetId, err := ioutil.ReadFile("sheetId")
	writeRange := "sheet1!A1:B1"
	value := [][]interface{}{{"IP Address", string(ipAddress)}}
	rb := &sheets.ValueRange{
		Range:  writeRange,
		Values: value,
	}

	valueInputOption := "RAW"

	fmt.Printf("Update Cells\n")
	_, err = srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, rb).ValueInputOption(valueInputOption).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("done!\n")
}
