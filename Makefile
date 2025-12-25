VERSION := $(shell git describe --tags --abbrev=0)
APP := nidavellir

run-web:
	cd web && npm run dev -- --host

run:
	go build . && ./$(APP)

test-seed:
	./nidavellir --cli --tokens --create --label test
	./nidavellir --cli --packages --register --name test --description test --repo test --type binary

build-release:
	GOOS=linux GOARCH=amd64 \
			 go build -tags release -ldflags "-X main.version=$(VERSION)" -o $(APP)
