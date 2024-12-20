# Overview ⚙️

Golang Plugin development + testing using the [`go-plugin`](https://github.com/hashicorp/go-plugin) package from [HashiCorp](https://github.com/hashicorp)

# Quick Start ⚡

```console
➜ git clone https://github.com/ohsu-comp-bio/plugins

➜ cd plugins

➜ ./build.sh

➜ ./authorize "Alyssa P. Hacker"
Before auth ➜ User: Alyssa P. Hacker
 After auth ➜ User: Alyssa P. Hacker, Token: <Alyssa's Secret Token>

➜ ./authorize "Eva Lu Ator"  
Before auth ➜ User: Eva Lu Ator
 After auth ➜ User: Eva Lu Ator, Token: <Eva's Secret Token>
```

# Additional Resources 📚

- https://github.com/hashicorp/go-plugin
- https://pkg.go.dev/github.com/hashicorp/go-plugin
- https://eli.thegreenplace.net/2023/rpc-based-plugins-in-go
- https://github.com/eliben/code-for-blog/tree/main/2023/go-plugin-htmlize-rpc
