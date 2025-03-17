.PHONY: all build

all: build

clean:
	rm kv-go-grpc
	rm kv

build:
	# Build main CLI 
	go build -o kv
	# Build Pluigin
	go build -o kv-go-grpc ./plugin
