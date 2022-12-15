GOARCH = 386
build: ./main.go
		go build main.go
		editbin /LARGEADDRESSAWARE ./main.exe