package server_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/maxcnunes/monitor-api/monitor/entities"
	. "github.com/maxcnunes/monitor-api/server"
)

var (
	router = Router{}
)

var _ = Describe("server", func() {
	Context("Performing GET request to '/targets' route", func() {
		BeforeEach(func() {
			dataMonitor.Target.Create("http://first-targe.com")
		})

		It("returns a 200 Status Code", func() {
			response := Request("GET", "/targets", nil, nil)
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns a list of targets", func() {
			var result []entities.Target
			_ = Request("GET", "/targets", &result, nil)

			Expect(len(result)).To(Equal(1))
			Expect(result[0].URL).To(Equal("http://first-targe.com"))
		})
	})

	Context("Performing POST request to '/targets' route", func() {
		It("returns a 201 Status Code", func() {
			target := entities.Target{URL: "http://second-targe.com"}
			response := Request("POST", "/targets", nil, target)
			Expect(response.Code).To(Equal(http.StatusCreated))
		})
	})

	Context("Performing PUT request to '/targets' route", func() {
		var target *entities.Target
		var result entities.Target

		BeforeEach(func() {
			target = dataMonitor.Target.Create("http://first-targe.com")

			target.URL = "http://updated-targe.com"
			target.Status = entities.StatusDown
		})

		It("returns a 200 Status Code", func() {
			response := Request("PUT", "/targets/"+target.ID.Hex(), &result, target)
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("updates the target in the database", func() {
			_ = Request("PUT", "/targets/"+target.ID.Hex(), &result, target)

			item := dataMonitor.Target.FindOneByID(target.ID.Hex())

			Expect(item.URL).To(Equal(target.URL))
			Expect(item.Status).To(Equal(target.Status))
		})
	})

	Context("Performing DELETE request to '/targets' route", func() {
		var target *entities.Target
		BeforeEach(func() {
			target = dataMonitor.Target.Create("http://first-targe.com")
		})

		It("returns a 200 Status Code", func() {
			response := Request("DELETE", "/targets/"+target.ID.Hex(), nil, nil)
			Expect(response.Code).To(Equal(http.StatusOK))
		})
	})

	Context("Performing GET request to '/tracks' route", func() {
		var target *entities.Target

		BeforeEach(func() {
			target = dataMonitor.Target.Create("http://first-targe.com")
			dataMonitor.Track.Create(*target, entities.StatusUp)
		})

		It("returns a 200 Status Code", func() {
			response := Request("GET", "/tracks", nil, nil)
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns a list of tracks", func() {
			var result []entities.Track
			_ = Request("GET", "/tracks", &result, nil)

			Expect(len(result)).To(Equal(1))
			Expect(result[0].TargetID).To(Equal(target.ID))
			Expect(result[0].Status).To(Equal(entities.StatusUp))
		})
	})
})
