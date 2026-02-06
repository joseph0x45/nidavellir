VERSION := $(shell git describe --tags --abbrev=0)
APP := gobase

pico.min.css:
	curl https://cdn.jsdelivr.net/npm/@picocss/pico@2.1.1/css/pico.min.css > pico.min.css

release:
	GOOS=linux GOARCH=amd64 \
		go build -tags release -ldflags "-X main.version=$(VERSION)" -o $(APP)
