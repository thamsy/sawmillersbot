# Saw Millers' Bot

A Telegram bot to inform the house of the duties they have.

## Setup

```bash
# After installing Go and setting up the GOPATH and GOROOT env variables
go get -u gopkg.in/telegram-bot-api.v4
go get -u google.golang.org/api/sheets/v4
go get -u golang.org/x/oauth2/...
```

### Telegram

Also, create a `secret` package containing the "Secret Bot Token" 

### Google Sheets API

In the same `secret` package, put the ID of the Google Sheets. Also, put the `client_secret.json` necessary for the sheetsapi oauth in the `GOPATH` directory. 

See Quickstart at https://developers.google.com/sheets/api/quickstart/go for more information.

## New Features

You may push any new features you may want in the issues tab.


