GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags '-w -s'  -o
BIN=bin/better-av-tool
BIN2=bin/better-av-tool.exe
SOURCE=.
VERSION=1.3.1

docker:
	$(GOBUILD) $(BIN) $(SOURCE)
macos:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(BIN) $(SOURCE)
	cp config.toml bin/
	zip bin/better-av-tool-darwin-amd64-v$(VERSION).zip $(BIN) bin/config.toml
win:
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(BIN2) $(SOURCE)
	cp config.toml bin/
	zip bin/better-av-tool-windows-amd64-v$(VERSION).zip $(BIN2) bin/config.toml
m1:
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(BIN) $(SOURCE)
	cp config.toml bin/
	zip bin/better-av-tool-darwin-arm64-v$(VERSION).zip $(BIN) bin/config.toml

