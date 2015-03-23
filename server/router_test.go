package server_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/maxcnunes/monitor-api/monitor"
	. "github.com/maxcnunes/monitor-api/server"
)

var (
	router = Router{}
)

var _ = Describe("server", func() {
	Context("Performing GET request to root route", func() {
		BeforeEach(func() {
			data.AddTarget("http://first-targe.com")
		})

		It("returns a 200 Status Code", func() {
			response := Request("GET", "/", router.ListHandler, nil, nil)
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns a list of targets", func() {
			var result []monitor.Target
			_ = Request("GET", "/", router.ListHandler, &result, nil)

			Expect(len(result)).To(Equal(1))
			Expect(result[0].URL).To(Equal("http://first-targe.com"))
		})
	})

	Context("Performing POST request to root route", func() {
		It("returns a 201 Status Code", func() {
			target := monitor.Target{URL: "http://second-targe.com"}
			response := Request("POST", "/", router.CreateHanler, nil, target)
			Expect(response.Code).To(Equal(http.StatusCreated))
		})
	})

	Context("Performing PUT request to root route", func() {
		var target *monitor.Target
		var result monitor.Target

		BeforeEach(func() {
			target = data.AddTarget("http://first-targe.com")

			target.URL = "http://updated-targe.com"
			target.Status = monitor.StatusDown
		})

		It("returns a 200 Status Code", func() {
			response := Request("PUT", "/"+target.ID.Hex(), router.UpdateHandler, &result, target)
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("updates the target in the database", func() {
			_ = Request("PUT", "/"+target.ID.Hex(), router.UpdateHandler, &result, target)

			item := data.GetTargetByID(target.ID.Hex())

			Expect(item.URL).To(Equal(target.URL))
			Expect(item.Status).To(Equal(target.Status))
		})
	})

	Context("Performing DELETE request to root route", func() {
		var target *monitor.Target
		BeforeEach(func() {
			target = data.AddTarget("http://first-targe.com")
		})

		It("returns a 200 Status Code", func() {
			response := Request("DELETE", "/"+target.ID.Hex(), router.DeleteHandler, nil, nil)
			Expect(response.Code).To(Equal(http.StatusOK))
		})
	})
})
