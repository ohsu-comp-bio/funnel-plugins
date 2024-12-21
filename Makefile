.PHONY: all build

all: build

build:
	@printf "Building ./authorize..."
	@mkdir -p plugin-binaries
	@go build -o ./plugin-binaries/exampleAuthorizer ./sample-plugins/exampleAuthorizer/
	@go build -o ./authorize ./main.go
	@echo "OK"
