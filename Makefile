format:
	go fmt ./...


build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs=false -trimpath -ldflags="-w -s" -o bin/forecast .

run: build
	bin/forecast