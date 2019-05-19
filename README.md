[![GoDoc](https://godoc.org/github.com/garsue/watermillzap?status.svg)](https://godoc.org/github.com/garsue/watermillzap)
[![CircleCI](https://circleci.com/gh/garsue/watermillzap.svg?style=svg)](https://circleci.com/gh/garsue/watermillzap)
[![Go Report Card](https://goreportcard.com/badge/github.com/garsue/watermillzap)](https://goreportcard.com/report/github.com/garsue/watermillzap)

# watermillzap

watermillzap provides the implementation of `watermill.LoggerAdapter` using [go.uber.org/zap](https://github.com/uber-go/zap).

## Notice

`Trace` writes debug log instead of trace log because zap does not support trace level logging.
