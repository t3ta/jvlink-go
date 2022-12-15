build: export GOOS=windows
build: export GOARCH=386
build:
	go build main.go
	editbin /LARGEADDRESSAWARE ./main.exe