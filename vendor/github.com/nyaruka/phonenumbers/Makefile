build:
	mkdir -p functions
	go get ./...
	go build -ldflags "-X main.Version=`git describe --tags`" -o functions/phoneserver ./cmd/phoneserver
