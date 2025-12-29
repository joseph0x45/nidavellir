VERSION := $(shell git describe --tags --abbrev=0)
APP := nidavellir

run-web:
	cd web && npm run dev -- --host

run:
	go build . && ./$(APP)

test-seed:
	./nidavellir --cli --tokens --create --label test
	./nidavellir --cli --tokens --create --label test2
	./nidavellir --cli --packages --register --name test --description test --repo test --type binary
	./nidavellir --cli --packages --register --name test2 --description test --repo test --type binary

release:
	cd web && npm install && npm run build
	GOOS=linux GOARCH=amd64 \
			 go build -tags release -ldflags "-X main.version=$(VERSION)" -o $(APP)
