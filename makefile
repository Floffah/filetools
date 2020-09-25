winbuild:
	go build -o bin/filetools.exe src/main.go

wintest:
	go build -o bin/filetools.exe src/main.go
	cd test && "../bin/filetools.exe" -Method=checkdir -Name=node_modules

unixbuild:
	go build -o bin/filetools src/main.go

debugbuild:
	go build -o bin/filetools -x src/main.go
