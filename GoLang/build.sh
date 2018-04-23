cd /workspace/2.GoLang/AutoUpdateDNS/
env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -o AutoUpdateDNS64.exe main.go
env CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=x86_64-w64-mingw32-gcc go build -o AutoUpdateDNS32.exe main.go
go build -o AutoUpdateDNSLinux main.go