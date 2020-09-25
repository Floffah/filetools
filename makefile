buildwin:
	go build -o bin/filetools.exe src/main.go

testwin

buildunix:
	go build -o bin/filetools src/main.go

builddebug:
	go build -o bin/filetools -x src/main.go
