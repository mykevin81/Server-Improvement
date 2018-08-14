# IP Address Upater
The purpose of this script is to update a home IP address to a google sheet to solve dynamic IP from ISP.
And also, I was bored.....

## Install
1. Make sure your GO is up to date, this program requires version 1.8+ to run correctly
2. Clone this repository in command line
```bash
git clone https://github.com/mykevin81/Server-Improvement.git
```
3. Setup the environment by following Step 1 & 2 of the quick start guide found [here](https://developers.google.com/sheets/api/quickstart/go), for step 2 please use the code below
```
go get -u google.golang.org/api/sheets/v4
go get -u golang.org/x/oauth2/...
```
4. Download `credentials.json` from the tutorial and put it in the same folder as `ipBroadcast.go`
5. Make your Google sheet in your Google drive and get the sheet ID from the URL
6. save the ID into a file and name it `sheetId` and put it in the same folder as `ipBroadcast.go`
7. Build the program by running `go build ipBroadcast.go`
8. Enjoy!

## Setup Cron Job to keep it updated
I set mine to every even hours, but you can always change up the frequency
```
0 * * * * ipBroadcast > ipBroadcast.log
```
For quick crontab generation, refer to this [guide](https://crontab-generator.org/)
