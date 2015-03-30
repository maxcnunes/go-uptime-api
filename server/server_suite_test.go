package server_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/maxcnunes/monitor-api/monitor/data"
)

var (
	db          = data.DB{}
	dataMonitor = data.DataMonitor{}
	handlers    *mux.Router
)

func TestMonitorApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Monitor-API-Server Suite")
}

// Setup server's tests
var _ = BeforeSuite(func() {
	db.Start()
	dataMonitor.Start(db)

	handlers = router.Start(&dataMonitor)
})

// Teardown server's tests
var _ = AfterSuite(func() {
	db.Close()
})

// Cleans the databse after each test
var _ = AfterEach(func() {
	db.Wipe()
})

//
// Helper functions for server's test
//

// Request sends a new request to the server been tested
func Request(method string, route string, result interface{}, body interface{}) *httptest.ResponseRecorder {
	bodyRequest := parseBodyRequest(body)

	request, _ := http.NewRequest(method, route, bodyRequest)
	response := httptest.NewRecorder()

	handlers.ServeHTTP(response, request)

	if result != nil {
		json.Unmarshal(response.Body.Bytes(), &result)
	}

	return response
}

func parseBodyRequest(item interface{}) io.Reader {
	if item == nil {
		return nil
	}

	body, err := json.Marshal(item)
	if err != nil {
		log.Println("Unable to marshal item")
	}

	return bytes.NewReader(body)
}
