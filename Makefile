build:
	@go build -o bin/billingGo

run: build
	@./bin/billingGo

test:
	@go test ./... -coverprofile=coverage