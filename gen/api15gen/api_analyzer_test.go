package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gopkg.in/rightscale/rsc.v5/gen"
)

var _ = Describe("APIAnalyzer ParseRoute", func() {
	var (
		moniker, route string

		pathPatterns []*gen.PathPattern
	)

	JustBeforeEach(func() {
		pathPatterns = ParseRoute(moniker, route)
	})

	Context("given a simple route", func() {
		BeforeEach(func() {
			route = "GET    /api/servers(.:format)? {:action=>\"index\", :controller=>\"servers\"}"
		})

		It("computes the path pattern", func() {
			Ω(len(pathPatterns)).Should(Equal(1))
			Ω(pathPatterns[0].HTTPMethod).Should(Equal("GET"))
			Ω(pathPatterns[0].Pattern).Should(Equal("/api/servers"))
			Ω(pathPatterns[0].Variables).Should(BeEmpty())
		})
	})

	Context("given an obsolete route", func() {
		BeforeEach(func() {
			route = "GET    /api/session(.:format)? {:action=>\"index\", :controller=>\"servers\"}"
		})

		It("does not produce a path pattern", func() {
			Ω(len(pathPatterns)).Should(Equal(0))
		})
	})

	Context("given a route with a variable", func() {
		BeforeEach(func() {
			route = "PUT    /api/servers/:id(.:format)? {:action=>\"index\", :controller=>\"servers\"}"
		})

		It("computes the path pattern", func() {
			Ω(len(pathPatterns)).Should(Equal(1))
			Ω(pathPatterns[0].HTTPMethod).Should(Equal("PUT"))
			Ω(pathPatterns[0].Pattern).Should(Equal("/api/servers/%s"))
			Ω(len(pathPatterns[0].Variables)).Should(Equal(1))
			Ω(pathPatterns[0].Variables[0]).Should(Equal("id"))
		})
	})

	Context("given a route with multiple variables", func() {
		BeforeEach(func() {
			route = "PUT    /api/clouds/:cloud_id/instances/:instance_id/security_groups/:id(.:format)? {:action=>\"index\", :controller=>\"security_groups\"}"
		})

		It("computes the path pattern", func() {
			Ω(len(pathPatterns)).Should(Equal(1))
			Ω(pathPatterns[0].HTTPMethod).Should(Equal("PUT"))
			Ω(pathPatterns[0].Pattern).Should(Equal("/api/clouds/%s/instances/%s/security_groups/%s"))
			Ω(len(pathPatterns[0].Variables)).Should(Equal(3))
			Ω(pathPatterns[0].Variables[0]).Should(Equal("cloud_id"))
			Ω(pathPatterns[0].Variables[1]).Should(Equal("instance_id"))
			Ω(pathPatterns[0].Variables[2]).Should(Equal("id"))
		})
	})

	Context("given multiple routes with multiple ", func() {
		BeforeEach(func() {
			route = "GET    /api/security_groups/:id(.:format)? {:action=>\"index\", :controller=>\"security_groups\"}"
			route += "GET    /api/instances/:instance_id/security_groups/:id(.:format)? {:action=>\"index\", :controller=>\"security_groups\"}"
			route += "GET    /api/clouds/:cloud_id/instances/:instance_id/security_groups/:id(.:format)? {:action=>\"index\", :controller=>\"security_groups\"}"
		})

		It("computes the path patterns", func() {
			Ω(len(pathPatterns)).Should(Equal(3))
			Ω(pathPatterns[0].HTTPMethod).Should(Equal("GET"))
			Ω(pathPatterns[0].Pattern).Should(Equal("/api/security_groups/%s"))
			Ω(len(pathPatterns[0].Variables)).Should(Equal(1))
			Ω(pathPatterns[0].Variables[0]).Should(Equal("id"))
			Ω(pathPatterns[0].Pattern).Should(Equal("/api/security_groups/%s"))
			Ω(len(pathPatterns[1].Variables)).Should(Equal(2))
			Ω(pathPatterns[1].HTTPMethod).Should(Equal("GET"))
			Ω(pathPatterns[1].Variables[0]).Should(Equal("instance_id"))
			Ω(pathPatterns[1].Variables[1]).Should(Equal("id"))
			Ω(pathPatterns[1].Pattern).Should(Equal("/api/instances/%s/security_groups/%s"))
			Ω(len(pathPatterns[2].Variables)).Should(Equal(3))
			Ω(pathPatterns[2].HTTPMethod).Should(Equal("GET"))
			Ω(pathPatterns[2].Variables[0]).Should(Equal("cloud_id"))
			Ω(pathPatterns[2].Variables[1]).Should(Equal("instance_id"))
			Ω(pathPatterns[2].Variables[2]).Should(Equal("id"))
			Ω(pathPatterns[2].Pattern).Should(Equal("/api/clouds/%s/instances/%s/security_groups/%s"))
		})
	})
})
