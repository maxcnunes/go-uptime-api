package server_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/maxcnunes/monitor-api/monitor"
	. "github.com/maxcnunes/monitor-api/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	db       = monitor.DB{}
	data     = monitor.DataMonitor{}
	router   = Router{}
	server   *httptest.Server
	usersURL string
)

func setup() {
	db.Start()
	data.Start(db)

	handlers := router.Start(&data)
	server = httptest.NewServer(handlers) //Creating new server with the user handlers

	usersURL = server.URL //Grab the address for the API endpoint
	fmt.Println(usersURL)
}

var _ = Describe("server", func() {

	BeforeEach(func() {
		setup()
		data.AddTarget("http://first-targe.com")
	})

	Context("List all targets", func() {
		It("returns a 200 Status Code", func() {
			request, _ := http.NewRequest("GET", "/", nil)
			response := httptest.NewRecorder()

			router.ListHandler(response, request)

			Expect(response.Code).To(Equal(http.StatusOK))
			// Request("GET", "/todos", router.ListHandler)
			// Expect(response.Code).To(Equal(200))
		})

		It("returns a list of targets", func() {
			request, _ := http.NewRequest("GET", "/", nil)
			response := httptest.NewRecorder()

			router.ListHandler(response, request)

			var targets []monitor.Target
			json.Unmarshal(response.Body.Bytes(), &targets)

			Expect(len(targets)).To(Equal(1))
			Expect(targets[0].URL).To(Equal("http://first-targe.com"))
		})
	})

	AfterEach(func() {
		db.Wipe()
		db.Close()
	})
})
