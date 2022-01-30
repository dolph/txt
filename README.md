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

Configure your Twilio account using environment variables supported by [twilio-go](https://github.com/twilio/twilio-go) (such as `TWILIO_ACCOUNT_SID` and `TWILIO_AUTH_TOKEN`, both of which are required).

`txt` directly consumes additional environment variables prefixed with `TXT_`:

* `TXT_FROM` (required): The Twilio phone number to use for sending.
* `TXT_TO` (required): The phone number to send a text message to.
