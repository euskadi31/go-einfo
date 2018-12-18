.PHONY: all clean test cover over-html

all: coverage.out

coverage.out: $(shell find . -type f -print | grep -v vendor | grep "\.go")
	@CGO_ENABLED=0 go test -cover -covermode=count -coverprofile ./coverage.out.tmp ./...
	@cat ./coverage.out.tmp | grep -v '.pb.go' | grep -v 'mock_' > ./coverage.out
	@rm ./coverage.out.tmp

test: coverage.out

cover: coverage.out
	@echo ""
	@go tool cover -func ./coverage.out

cover-html: coverage.out
	@go tool cover -html=./coverage.out

lint:
	@golangci-lint run ./...

clean:
	@rm ./coverage.out
	@go clean -i ./...
