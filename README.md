logwrapper
==========

[![CI](https://github.com/nekrassov01/logwrapper/actions/workflows/test.yml/badge.svg)](https://github.com/nekrassov01/logwrapper/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/nekrassov01/logwrapper/graph/badge.svg)](https://codecov.io/gh/nekrassov01/logwrapper)
[![Go Reference](https://pkg.go.dev/badge/github.com/nekrassov01/logwrapper.svg)](https://pkg.go.dev/github.com/nekrassov01/logwrapper)
[![Go Report Card](https://goreportcard.com/badge/github.com/nekrassov01/logwrapper)](https://goreportcard.com/report/github.com/nekrassov01/logwrapper)

Small, thin wrapper for logging in Go and AWS SDK

Example
-------

```go
var logger *log.AppLogger

func main() {
	// Create global logger for the application
	l, err := log.NewAppLogger(os.Stderr, "debug", "default", "MyApp")
	if err != nil {
		panic(err)
	}
	logger = l

	...

	// Create logger for AWS SDK called through the application
	s, err := log.NewSDKLogger(os.Stderr, "debug", "default", "SDK")
	if err != nil {
		panic(err)
	}
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithLogger(s))
	if err != nil {
		panic(err)
	}

	...
}

```

Using
-----

Wraps the following:

- [github.com/charmbracelet/log](https://github.com/charmbracelet/log)
- [github.com/aws/smithy-go](https://github.com/aws/smithy-go)

Author
------

[nekrassov01](https://github.com/nekrassov01)

License
-------

[MIT](https://github.com/nekrassov01/logwrapper/blob/main/LICENSE)
