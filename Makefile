.PHONY: all build clean

all: build

clean:
	# Remove main CLI + Plugin
	rm kv
	rm authorize-go-grpc

build:
	# Build main CLI 
	go build -o authorize

	# Build Pluigin
	go build -o authorize-go-grpc ./plugin
