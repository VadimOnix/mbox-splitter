# run project
run:
	go run src/main/main.go src/assets/sample.mbox

build:
	GOOS=linux GOARCH=amd64 go build -o dist/mboxsplit_linux_x86_64 src/main/main.go
	GOOS=windows GOARCH=amd64 go build -o dist/mboxsplit_win_x86_64.exe src/main/main.go
	GOOS=darwin GOARCH=arm64 go build -o dist/mboxsplit_darwin_arm64 src/main/main.go