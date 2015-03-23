package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/maxcnunes/monitor-api/monitor"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var (
	db       = monitor.DB{}
	data     = monitor.DataMonitor{}
)

func TestMonitorApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Monitor-API-Server Suite")
}

// Setup server's tests
var _ = BeforeSuite(func() {
	db.Start()
	data.Start(db)

	handlers := router.Start(&data)
	server = httptest.NewServer(handlers)
})

// Teardown server's tests
var _ = AfterSuite(func() {
	db.Wipe()
	db.Close()
})

//
// Helper functions for server's test
//

// Request sends a new request to the server been tested
func Request(method string, route string, handler http.HandlerFunc, result interface{}) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, route, nil)
	response := httptest.NewRecorder()

	handler(response, request)

	if result != nil {
		json.Unmarshal(response.Body.Bytes(), &result)
	}

	return response
}
