# txt

Send text messages from the command line.

## Usage

Configuration is provided by environment variables. All arguments are sent as a text message:

```bash
$ txt hello world
```

To specify the recipient, use `TXT_TO`:

```bash
$ TXT_TO=15555551234 txt hello world
```

### Configuration

`txt` consumes environment variables prefixed with `TXT_`, including:

* `TXT_TWILIO_ACCOUNT_SID` (required): Your Twilio account string identifier.
* `TXT_TWILIO_AUTH_TOKEN` (required): Your Twilio account auth token.
* `TXT_TWILIO_FROM` (required): The Twilio phone number to use for sending.
* `TXT_TO` (required): The phone number to send a text message to.
