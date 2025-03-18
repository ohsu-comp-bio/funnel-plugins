.PHONY: all build clean

all: build test-server

clean:
	# Remove CLI
	rm -f authorizer

	# Remove Plugin
	rm -f authorizer-plugin
	
	# Remove test server
	rm -f test-server

build:
	# Build CLI 
	go build -o authorizer

	# Build Plugin
	go build -o authorizer-plugin ./plugin

test-server:
	# Build test server
	go build -o test-server ./tests/test-server
