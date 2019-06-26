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
	logger = log.New(os.Stdout, "dhclient: ", log.LstdFlags)

	helpFlag  = flag.BoolP("help", "h", false, "display this help dialog")
	healthUrl = flag.String("url", "http://localhost/", "health endpoint to check")
	jsonExpr  = flag.String("path", "", "optional json path to check")
	value     = flag.String("value", "", "json value to check")
)

func main() {
	flag.Parse()
	if *helpFlag {
		help()
		os.Exit(0)
	}

	/*
		see https://docs.docker.com/engine/reference/builder/#healthcheck
		The command’s exit status indicates the health status of the container. The possible values are:

		0: success - the container is healthy and ready for use
		1: unhealthy - the container is not working correctly
		2: reserved - do not use this exit code
	*/
	logger.Printf("checking url %s", *healthUrl)
	resp, err := http.Get(*healthUrl)
	if err != nil {
		panic(fmt.Errorf("get failed on url with error: %v", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		if *jsonExpr != "" {
			op, err := jq.Parse(*jsonExpr)
			if err != nil {
				panic(fmt.Errorf("failed to parse JSON expression: %v", err))
			}
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(fmt.Errorf("failed to read response body: %v", err))
			}
			data, _ := op.Apply(bodyBytes)
			if string(data) == *value {
				os.Exit(0)
			}
		} else {
			os.Exit(0)
		}
	}

	// otherwise
	os.Exit(1)
}

var helpMsg = `dhclient - health check a URL
usage: dhclient`

func help() {
	fmt.Println(helpMsg)
	flag.PrintDefaults()
}
