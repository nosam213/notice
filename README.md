# notice
notice is designed to read ssrf requests headers and forms 

### Flags
* -h, -help
* -from-one			starts request counting from 1
* -hide-banner		disables info banner
* -ip				the ip (default "0.0.0.0")
* -local-time		enable local time timestamps (default UTC)
* -port				the port (default "9001")
* -tls				TLS (default @ ./cert.pem ./key.pem)

### TLS/SSL
NOTE: If not specificed, requires exactly named 'cert.pem' and 'key.pem' in current directory.
#### Quick Certificate
`openssl req -nodes -newkey rsa:4096 -sha512 -new -x509 -days 3650 -keyout key.pem -out cert.pem`
```
$ ls
cert.pem  key.pem
$ notice -tls --port 8443
notice running at: https://0.0.0.0:9001 [version: 0.1.1]
```

## Compiling
NOTE: Requires network connection for the first compile.

### Windows
```
C:\Users\username\notice> dir
go.mod  go.sum  main.go  README.md
C:\Users\username\notice> go build -o notice.exe -ldflags="-w -s -buildid=" .
go.mod  go.sum  notice.exe  main.go  README.md
```

### Unix
```
$ export CGO_ENABLED=0 // Enables static compiling (optional)
$ pwd
/home/username/notice
$ ls
go.mod  go.sum  main.go  README.md
$ go build -o notice -ldflags="-w -s -buildid=" .
go.mod  go.sum  notice  main.go  README.md
```
