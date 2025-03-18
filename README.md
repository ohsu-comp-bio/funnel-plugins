[![Go Build + Test](https://github.com/ohsu-comp-bio/funnel-plugins/actions/workflows/tests.yaml/badge.svg)](https://github.com/ohsu-comp-bio/funnel-plugins/actions/workflows/tesrs.yaml)
[![Project license](https://img.shields.io/github/license/ohsu-comp-bio/funnel-plugins.svg)](LICENSE)
[![Coded with love by ohsu-comp-bio](https://img.shields.io/badge/Coded%20with%20%E2%99%A5%20by-OHSU-blue)](https://github.com/ohsu-comp-bio)

</div>

# Overview ⚙️

> [!NOTE]
> Adapted from Hashicorp's [gRPC KV Plugin example](https://github.com/hashicorp/go-plugin/tree/main/examples/grpc) 🚀

This repo contains Funnel Plugin development using the [`go-plugin`](https://github.com/hashicorp/go-plugin) package from [HashiCorp](https://github.com/hashicorp).

In this setup, the Plugin handles all user authentication, with the Server having no knowledge or record of specific user credentials/tokens.

# Quick Start ⚡

## 1. Start the Test Server 

```console
➜ git clone https://github.com/ohsu-comp-bio/funnel-plugins

➜ cd funnel-plugins

➜ make test-server

➜ ./test-server
Server is running on http://localhost:8080
```

## 2. Build the `authorizer` Plugin

```sh
➜ make

➜ export FUNNEL_PLUGIN=./authorizer-plugin
```

## 3. Get Authorized User ✅

Here we invoke the CLI component request to authenticate a user named `example` who is an `Authorized` user (i.e. found in the "User Database" — [`example-users.csv`](./tests/example-users.csv)):

```sh
➜ ./authorizer example | jq
{
  "token": "example's secret",
  "user": "example"
}
```

## 4. Get Unauthorized User ❌

Here we attempt to authenticate a user named `error`, representing an `Unauthorized` user:

```sh
➜ ./authorizer error | jq
{
  "error": "user 'error' not found"
}
```

# Architecture 📐

This repo contains the following major components:
1. CLI (`authorizer-cli`) — used to manually invoke the `authorizer` plugin from the command line
2. Plugin (`authorizer`) — the actual plugin itself that makes a call to the "external" Test Server
3. Test Server (`test-server`) — used as a mock service to store the actual users and their tokens/credentials 

# Sequence Diagram 📝

> Created with https://sequencediagram.org ([_source_](https://github.com/ohsu-comp-bio/funnel/blob/feature/plugins/plugins/sequence-diagram.txt))

![proposed-auth-design](./sequence-diagram.png)

# Additional Resources 📚

- https://github.com/hashicorp/go-plugin
- https://pkg.go.dev/github.com/hashicorp/go-plugin
- https://eli.thegreenplace.net/2023/rpc-based-plugins-in-go
- https://github.com/eliben/code-for-blog/tree/main/2023/go-plugin-htmlize-rpc
