DIR = build

.PHONY: all build clean tests

all: build tests

clean:
	rm -rf $(DIR)/*

build:
	# Build CLI 
	go build -o $(DIR)/cli

	# Build Plugin
	go build -o $(DIR)/plugins/authorizer ./plugin

tests:
	# Build test server
	go build -o $(DIR)/test-server ./tests/test-server.go
