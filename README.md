# remote console for web browser

## usage
server
```bash
$ go run server.go
```

client
```bash
$ curl localhost:5500/\?type=log\&msg=hogehoge
```
type: (log|info|warn|err)
