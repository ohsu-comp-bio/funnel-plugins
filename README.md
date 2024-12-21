> Adapted from [*RPC-based plugins in Go*](https://eli.thegreenplace.net/2023/rpc-based-plugins-in-go) by [Eli Bendersky](https://eli.thegreenplace.net/) 🚀

# Overview ⚙️

Golang Plugin development + testing using the [`go-plugin`](https://github.com/hashicorp/go-plugin) package from [HashiCorp](https://github.com/hashicorp)

# Quick Start ⚡

## 1. Start the Server 
First start the server:

```console
➜ git clone https://github.com/ohsu-comp-bio/plugins

➜ cd plugins

➜ make
Building ./authorize...OK

➜ ./authorize
Listening on http://localhost:8080
```

## 2. Send Requests

In another terminal, send the request using of the example users below:

```console
➜ curl -H "Authorization: Bearer Alyssa P. Hacker" http://localhost:8080

User: Alyssa P. Hacker, Token: <Alyssa's Secret> ✅
```

Here's an example of a user that's not found by the plugin (representing an `Unauthorized` user):

```console
➜ curl -H "Authorization: Bearer Foobar" http://localhost:8080

Error: User Foobar not found ❌
```

# Example Users ✍️

[Example users](https://en.wikipedia.org/wiki/Structure_and_Interpretation_of_Computer_Programs#Characters) may be found in [example-users.csv](./example-users.csv):
> - Alyssa P. Hacker, a Lisp hacker
> - Ben Bitdiddle
> - Cy D. Fect, a "reformed C programmer"
> - Eva Lu Ator
> - Lem E. Tweakit
> - Louis Reasoner, a loose reasoner

# Additional Resources 📚

- https://github.com/hashicorp/go-plugin
- https://pkg.go.dev/github.com/hashicorp/go-plugin
- https://eli.thegreenplace.net/2023/rpc-based-plugins-in-go
- https://github.com/eliben/code-for-blog/tree/main/2023/go-plugin-htmlize-rpc
