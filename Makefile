build-mac:
	GOOS=darwin GOARCH=arm64 go build -o ./bin/darwin-arm64/SettingsService -ldflags "-s -w -X main.Version=$(version)" ./cmd/SettingsService
	
build-lin:	
	GOOS=linux GOARCH=386 go build -o ./bin/linux-386/SettingsService -ldflags "-s -w -X main.Version=$(version)" ./cmd/SettingsService
	
build-win:	
	GOOS=windows GOARCH=386 go build -o ./bin/windows-386/SettingsService.exe -ldflags "-s -w -X main.Version=$(version)" ./cmd/SettingsService