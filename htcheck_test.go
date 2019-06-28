/*
  test for htcheck using command line parameters
*/
package main

import (
	"github.com/jarcoal/httpmock"
	"gotest.tools/assert"
	"testing"
)

func TestURLOnlyOK(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/url-ok",
		httpmock.NewStringResponder(200, ``))

	testUrl(t, 0, "http://localhost/url-ok")

	assert.Equal(t, httpmock.GetTotalCallCount(), 1, "mock http server not called")
}

func TestURLOnlyFail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/url-fail",
		httpmock.NewStringResponder(404, ``))

	testUrl(t, 1, "http://localhost/url-fail")

	assert.Equal(t, httpmock.GetTotalCallCount(), 1, "mock http server not called")
}

func TestJSONOK(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/json-up",
		httpmock.NewStringResponder(200, `{"status": "UP"}`))

	testJSON(t, 0, "http://localhost/json-up", ".status", "UP")

	assert.Equal(t, httpmock.GetTotalCallCount(), 1, "mock http server not called")
}

func TestJSONFail(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "http://localhost/json-down",
		httpmock.NewStringResponder(200, `{"status": "DOWN"}`))
	httpmock.GetTotalCallCount()

	testJSON(t, 1, "http://localhost/json-down", ".status", "UP")

	assert.Equal(t, httpmock.GetTotalCallCount(), 1, "mock http server not called")
}

func testUrl(t *testing.T, expectedExitCode int, url string) {
	exitCode := checkUrl(url, "", "")
	assert.Equal(t, expectedExitCode, exitCode, "did not get expected exit code")
}
func testJSON(t *testing.T, expectedExitCode int, url string, path string, value string) {
	exitCode := checkUrl(url, path, value)
	assert.Equal(t, expectedExitCode, exitCode, "did not get expected exit code")
}
