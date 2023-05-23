#!/bin/bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/denter-darwin-amd64 .
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/denter-darwin-arm64 .

CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o build/denter-windows-386 .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/denter-windows-amd64 .

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/denter-linux-amd64 .
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o build/denter-linux-386 .
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o build/denter-linux-arm .
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/denter-linux-arm64 .
CGO_ENABLED=0 GOOS=linux GOARCH=mips go build -o build/denter-linux-mips .
CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build -o build/denter-linux-mips64 .
