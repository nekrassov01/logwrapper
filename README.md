logwrapper
==========

[![CI](https://github.com/nekrassov01/logwrapper/actions/workflows/test.yml/badge.svg)](https://github.com/nekrassov01/logwrapper/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/nekrassov01/logwrapper/graph/badge.svg)](https://codecov.io/gh/nekrassov01/logwrapper)
[![Go Reference](https://pkg.go.dev/badge/github.com/nekrassov01/logwrapper.svg)](https://pkg.go.dev/github.com/nekrassov01/logwrapper)
[![Go Report Card](https://goreportcard.com/badge/github.com/nekrassov01/logwrapper)](https://goreportcard.com/report/github.com/nekrassov01/logwrapper)

Small, thin wrapper for logging with Go and AWS SDK

Example
-------

```go
// Parse log level
level, err := log.ParseLevel("debug")
if err != nil {
	return err
}

// Parse log styles
styles, err := log.ParseStyles("labeled")
if err != nil {
	return err
}

// Create logger for the application
appLogger := log.NewAppLogger(os.Stderr, level, styles, "MyApp")

// Create logger for AWS SDK called through the application
sdkLogger := log.NewSDKLogger(os.Stderr, level, styles, "SDK")
cfg, err := config.LoadDefaultConfig(context.Background(), config.WithLogger(sdkLogger))
if err != nil {
	return err
}

...
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
