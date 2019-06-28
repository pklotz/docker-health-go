# docker-health-go (htcheck)
Special-purpose HTTP client for light-weight health check for use in docker

Why this program?
When using docker directly or via docker-compose, you can and should define a health check, so that docker 
knows that the process it is running is doing well. There are a couple of libraries to provide a HTTP health endpoints
for Go, such as [https://github.com/docker/go-healthcheck] and also Java offers with Spring Boot Actuator corresponding
frameworks. 

But on the client side, you still need to use curl or the outdated wget to perform the check.
If you ever checked, which dependencies curl and thus libcurl4 brings with it, you might wonder if this is worth 
the balast just to do a simple HTTP get with an exit code. Libcurl brings openldap libraries into the image 
and what not. So this little decent project provides a special-purpose HTTP client to use for health checks
in docker or elsewhere instead of throwing a general-purpose HTTP client at the job.

## Build

Build binary and test:
```
go build -o htcheck
go test
```

when you would like to get rid of a bit fat from the statically linked binary, use the flags
```
go build -o htcheck -ldflags="-s -w"
```

## Usage

Simple sample usage in a Dockerfile:
```
COPY ./htcheck /usr/bin/

HEALTHCHECK --interval=5m --timeout=3s CMD htcheck -u http://localhost/ || exit 1
```

Sample usage for Spring Boot actuator health endpoint:
```
COPY ./htcheck /usr/bin/

HEALTHCHECK --interval=5m --timeout=3s CMD htcheck -u http://localhost/health -p .status -v UP || exit 1
```

## Internals

Dependencies used by the project:
- pflag: posix-compatible command line flags (https://github.com/spf13/pflag) - BSD 3-clause license
- jq: jq-style JSON path (https://github.com/savaki/jq) - Apache 2 license

Testing only:
- httpmock: HTTP server mocking test library (https://github.com/jarcoal/httpmock) - MIT license
- gotesttools: asserts for go (https://godoc.org/gotest.tools) - Apache 2 license