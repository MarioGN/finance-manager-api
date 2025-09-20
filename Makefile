build:
	@go build -o bin/financemanager-api

run: build
	@./bin/financemanager-api

test:
	@go test -v ./...

test-cover:
	@mkdir -p coverage-report
	@go test -coverprofile=coverage-report/coverage.out ./...
	@go tool cover -html=coverage-report/coverage.out -o coverage-report/coverage.html
