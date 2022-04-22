# ![HAProxy](assets/images/haproxy-weblogo-210x49.png "HAProxy")

## Go Logger

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
[![Github tag](https://badgen.net/github/tag/haproxytech/go-logger)](https://github.com/haproxytech/go-logger/tags/)
[![CI](https://github.com/haproxytech/go-logger/actions/workflows/actions.yaml/badge.svg)](https://github.com/haproxytech/go-logger/actions/workflows/actions.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/haproxytech/go-logger)](https://goreportcard.com/report/github.com/haproxytech/go-logger)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/haproxytech/go-logger)
[![go.mod Go version](https://img.shields.io/github/go-mod/go-version/haproxytech/go-logger.svg)](https://github.com/haproxytech/go-logger)
[![Contributors](https://img.shields.io/github/contributors/haproxytech/go-logger?color=purple)](https://github.com/haproxy/haproxy/blob/master/CONTRIBUTING)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)





### Description

Go Logger is an interface for logging in Go programing language.
Beside interface it also provides option to create a logger that can log to multiple destinations

### Using multiple loggers at the same time

```go
import (
  ...
  logger "github.com/haproxytech/go-logger"
  ...
)

log1 := // create log1
log2 := // create log2

log := logger.New(log1, log2)

log.Print("Printing to two different loggers")
```

### Contributing

Thanks for your interest in the project and your willing to contribute:

- Pull requests are welcome!
- For commit messages and general style please follow the haproxy project's [CONTRIBUTING guide](https://github.com/haproxy/haproxy/blob/master/CONTRIBUTING) and use that where applicable.
- Please use `golangci-lint run` from [github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint) for linting code.

### Discussion

A Github issue is the right place to discuss feature requests, bug reports or any other subject that needs tracking.

To ask questions, get some help or even have a little chat, you can join our #ingress-controller channel in [HAProxy Community Slack](https://slack.haproxy.org).

## License

[Apache License 2.0](LICENSE)
