This app will do a simple task. It translate normal text to Base64.
You can write a config file, encode it in your console and paste in k8s secrets

usage: 

```Use it to decode/encode base64 string or files

## Usage:
  base64 [command]

Available Commands:
  decode      Decode file from base64
  encode      Encode file to base64
  help        Help about any command

## Flags:
  -h, --help   help for base64```

## Example:

### Decode:

```gobase64 decode MTIzNDU2Nzg5
123456789```


### Encode:

``` gobase64 encode alertmanager.yaml
Imdsb2JhbCI6CiAgInJlc29sdmVfdGltZW91dCI6ICI1bSIKInJlY2VpdmVycyI6Ci0gIm5hbWUiOiAibnVsbCIKInJvdXRlIjoKICAiZ3JvdXBfYnkiOgogIC0gImpvYiIKICAiZ3JvdXBfaW50ZXJ2YWwiOiAiNW0iCiAgImdyb3VwX3dhaXQiOiAiMzBzIgogICJyZWNlaXZlciI6ICJudWxsIgogICJyZXBlYXRfaW50ZXJ2YWwiOiAiMTJoIgogICJyb3V0ZXMiOgogIC0gIm1hdGNoIjoKICAgICAgImFsZXJ0bmFtZSI6ICJXYXRjaGRvZyIKICAgICJyZWNlaXZlciI6ICJudWxsIg== ```
