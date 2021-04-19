doc:
	swag init -o ./api

build:
	CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s -w' -o blogservice .

test:
	go test ./...

start:
	./blogservice --config=./configs

format:
	gofmt -w .

dev:
	go run main.go

lint:
	golangci-lint run -E golint -E lll -E goconst -E gosec -E gofmt ./...
