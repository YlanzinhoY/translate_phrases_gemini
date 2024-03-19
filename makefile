run:
	go run main.go

build:
	GOOS=linux GOARCH=amd64 go build -o build/translate_phrase_linux
	GOOS=darwin GOARCH=amd64 go build -o build/translate_phrase_macos
	GOOS=windows GOARCH=amd64 go build -o build/translate_phrase.exe