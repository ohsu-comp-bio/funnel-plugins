.PHONY: all build

all: build

clean:
	rm ./plugin/authorizer
	rm basic

build:
	go build -o ./plugin/authorizer ./plugin/auth_impl.go
	go build -o basic . 
