# IP Address Upater
The purpose of this script is to update a home IP address to a google sheet to solve dynamic IP from ISP.
And also, I was bored.....

## Install
1. Make sure your GO is up to date, this program requires version 1.8+ to run correctly
2. Clone this repository
3. Setup the environment by following Step 1 & 2 of the quick start guide found [here](https://developers.google.com/sheets/api/quickstart/go)
4. Init the go project to get the dependencies according to the init output
```
go mod init ipBroadcast.go
go mod tidy
```
5. Download `credentials.json` from the tutorial and put it in the same folder as `ipBroadcast.go`
6. Make your Google sheet in your Google drive and get the sheet ID from the URL
7. save the ID into a file and name it `sheetId` and put it in the same folder as `ipBroadcast.go`
8. Build the program by running `go build ipBroadcast.go`
9. Enjoy!

## Setup Cron Job to keep it updated
I set mine to every even hours, but you can always change up the frequency
```
0 * * * * ipBroadcast > ipBroadcast.log
```
For quick crontab generation, refer to this [guide](https://crontab-generator.org/)
