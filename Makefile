git_commit := $(shell git rev-parse --short HEAD)
git_branch := $(shell git symbolic-ref -q --short HEAD)
git_upstream := $(shell git remote get-url $(shell git config branch.$(shell git symbolic-ref -q --short HEAD).remote) 2> /dev/null)
export GIT_BRANCH = $(git_branch)
export GIT_UPSTREAM = $(git_upstream)

VERSION_LDFLAGS=\
 -X "github.com/ohsu-comp-bio/funnel/version.BuildDate=$(shell date)" \
 -X "github.com/ohsu-comp-bio/funnel/version.GitCommit= $(git_commit)" \
 -X "github.com/ohsu-comp-bio/funnel/version.GitBranch=$(git_branch)" \
 -X "github.com/ohsu-comp-bio/funnel/version.GitUpstream=$(git_upstream)"

DIR = build

.PHONY: all build clean tests

all: build tests

clean:
	rm -rf $(DIR)/*

build:
	# Build CLI 
	go build -o $(DIR)/cli

	# Build Plugin
	go build -o $(DIR)/plugins/authorizer ./plugin-go

tests:
	# Build test server
	go build -o $(DIR)/test-server ./tests/test-server.go

# Build binaries for all OS/Architectures
snapshot: release-dep
	@goreleaser \
		--clean \
		--snapshot

# Create a release on Github using GoReleaser
release:
	@goreleaser --clean

# Install dependencies for release
release-dep:
	@go get github.com/goreleaser/goreleaser
	@go get github.com/buchanae/github-release-notes
