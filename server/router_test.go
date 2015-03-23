package server_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/maxcnunes/monitor-api/monitor"
	. "github.com/maxcnunes/monitor-api/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	router   = Router{}
	server   *httptest.Server
	usersURL string
)

var _ = Describe("server", func() {
	BeforeEach(func() {
		data.AddTarget("http://first-targe.com")
	})

	Context("List all targets", func() {
		It("returns a 200 Status Code", func() {
			response := Request("GET", "/", router.ListHandler, nil)
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns a list of targets", func() {
			var targets []monitor.Target
			_ = Request("GET", "/", router.ListHandler, &targets)

			Expect(len(targets)).To(Equal(1))
			Expect(targets[0].URL).To(Equal("http://first-targe.com"))
		})
	})
})
