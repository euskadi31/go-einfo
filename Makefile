
deps:
	@go get github.com/mattn/go-isatty
	@go get github.com/stretchr/testify/assert

test: deps
	@go test ./...

cov: deps
	@go test ./... -cover
