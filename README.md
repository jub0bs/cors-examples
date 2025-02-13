# cors-examples

## What is this?

This repository is a collection of examples that illustrate how to use
[github.com/jub0bs/cors][jub0bs-cors], a principled CORS middleware library,
in conjunction with [`net/http`][net-http] and popular third-party routers
for [Go][go].

## What is the example?

The example invariably consists in creating a CORS middleware that

- allows anonymous access from Web origin `https://example.com`,
- with requests whose method is either `GET` or `POST`, and
- (optionally) with request header `Authorization`,

and applying the middleware in question to all the resources accessible
under an `/api/` path.

In addition, a `/hello` path is left unconfigured for CORS.

## Which routers are represented?

| Library              | Compatibility | Path to example                            |
| -------------------- | ------------- | ------------------------------------------ |
| [Chi][chi]           | v5.2+         | [chi=v5/main.go](chi-v5/main.go)           |
| [Echo][echo]         | v4.13+        | [echo-v4/main.go](echo-v4/main.go)         |
| [Fiber][fiber]       | v2.52+        | [fiber-v2/main.go](fiber-v2/main.go)       |
| [gorilla/mux][mux]   | v1.8+         | [gorilla-mux/main.go](gorilla-mux/main.go) |
| [net/http][net-http] | v1.22+        | [net-http/main.go](net-http/main.go)       |

[chi]: https://go-chi.io/#/
[echo]: https://echo.labstack.com/
[fiber]: https://gofiber.io/
[go]: https://go.dev/
[jub0bs-cors]: https://pkg.go.dev/github.com/jub0bs/cors
[net-http]: https://pkg.go.dev/net/http
[mux]: https://gorilla.github.io/
