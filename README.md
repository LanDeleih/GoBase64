This app doing a simple task to translate text or file into Base64.

Build:<br>

``` make build```

Use it to decode/encode base64 string or files:

```
NAME:
   goBase64

USAGE:
    [global options] command [command options] [arguments...]

VERSION:
   dev

DESCRIPTION:
   decode and encode plain text or file

COMMANDS:
   decode, d  Decode base64 file or string
   encode, e  Encode base64 file or string
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```

#### Examples:

Encode string:<br>
```goBase64 encode Hello-World```

Decode:<br>
```goBase64 decode SGVsbG8tV29ybGQ=```

Encode file:<br>
```goBase64 encode -f example/alertmanager.yaml```


