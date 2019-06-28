# docker-health-go
Special-purpose HTTP client for light-weight health check for use in docker

# htcheck Internals

Dependencies used by the project:
- pflag: posix-compatible command line flags (github.com/spf13/pflag) - BSD 3-clause license
- jq: jq-style JSON path (github.com/savaki/jq) - Apache 2 license

Testing only:
- httpmock: HTTP server mocking test library (https://github.com/jarcoal/httpmock) - MIT license
- gotesttools: asserts for go (https://godoc.org/gotest.tools) - Apache 2 license