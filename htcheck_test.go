package main

import (
	"github.com/jarcoal/httpmock"
	"os/exec"
	"testing"
)

func testURLOnlyOK(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "https://localhost/url-ok",
		httpmock.NewStringResponder(200, ``))

	cmd := exec.Command("dhclient", "--url", "https://localhost/url-ok")
	err := cmd.Run()
	if err != nil {
		t.Fail()
	}
}

func testURLOnlyFail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "https://localhost/url-fail",
		httpmock.NewStringResponder(404, ``))
}

func testJSONOK(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "https://localhost/json-up",
		httpmock.NewStringResponder(200, `{"status": "UP"}`))
}

func testJSONFail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "https://localhost/json-down",
		httpmock.NewStringResponder(200, `{"status": "down"}`))
}
