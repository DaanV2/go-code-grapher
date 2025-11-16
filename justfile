default:
	just --list

documentation:
	go doc -all -http

build:
	go build ./...

test:
	go test ./... --cover -coverprofile=reports/coverage.out --covermode atomic --coverpkg=./...

show-coverage-report:
	go tool cover -html=reports/coverage.out

coverage-report: test show-coverage-report

generate:
	go generate ./...

lint:
	go tool golangci-lint run -v --fix

format:
	go fmt ./...