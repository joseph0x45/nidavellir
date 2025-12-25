VERSION := $(shell git describe --tags --abbrev=0)
APP := nidavellir

run-web:
	cd web && npm run dev -- --host

run:
	go build . && ./$(APP)

build-release:
	GOOS=linux GOARCH=amd64 \
			 go build -tags release -ldflags "-X main.version=$(VERSION)" -o $(APP)
