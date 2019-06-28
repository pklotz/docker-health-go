package main

import (
	"fmt"
	"github.com/savaki/jq"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	logger = log.New(os.Stdout, "htcheck: ", log.LstdFlags)

	helpFlag  = flag.BoolP("help", "h", false, "display this help dialog")
	healthUrl = flag.StringP("url", "u", "http://localhost/", "health endpoint to check")
	jsonExpr  = flag.StringP("path", "p", "", "optional json path to check")
	value     = flag.StringP("value", "v", "", "json value to check")
)

func main() {
	flag.Parse()
	if *helpFlag {
		help()
		os.Exit(0)
	}

	os.Exit(checkUrl(*healthUrl, *jsonExpr, *value))
}

func checkUrl(healthUrl string, path string, value string) int {
	/*
		see https://docs.docker.com/engine/reference/builder/#healthcheck
		The command’s exit status indicates the health status of the container. The possible values are:

		0: success - the container is healthy and ready for use
		1: unhealthy - the container is not working correctly
		2: reserved - do not use this exit code
	*/
	logger.Printf("checking url %s", healthUrl)
	resp, err := http.Get(healthUrl)
	if err != nil {
		logger.Printf("get failed on url with error: %v", err)
		return 1
	}
	defer resp.Body.Close()
	logger.Printf("response status code: %d", resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		if path != "" {
			op, err := jq.Parse(path)
			if err != nil {
				logger.Printf("failed to parse JSON expression: %v", err)
				return 2
			}
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logger.Printf("failed to read response body: %v", err)
				return 2
			}
			data, _ := op.Apply(bodyBytes)
			logger.Printf("JSON value received is %s", data)
			if value == trimQuotes(string(data)) {
				return 0
			}
		} else {
			// response code not ok
			return 0
		}
	}
	// otherwise
	return 1
}

func trimQuotes(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

var helpMsg = `dhclient - health check a URL
usage: dhclient`

func help() {
	fmt.Println(helpMsg)
	flag.PrintDefaults()
}
