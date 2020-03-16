This app will do a simple task to translate text into Base64.

Build:<br>

``` make build```

Use it to decode/encode base64 string or files:

```
goBase64 [command]

Available Commands:
  decode      Decode file from base64
  encode      Encode file to base64
  help        Help about any command
```
#### Flags:
  -h, --help   help for base64

#### Examples:

Decode:<br>
```./goBase64 decode SGVsbG8tV29ybGQ=```

Encode file:<br>
```goBase64 encode -f example/alertmanager.yaml```

Encode string:<br>
```goBase64 encode SGVsbG8tV29ybGQ=```
