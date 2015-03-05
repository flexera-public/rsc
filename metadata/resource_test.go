package metadata_test

import (
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rightscale/rsc/metadata"
)

var _ = Describe("Resource", func() {

	Context("ExtractVariables", func() {
		var (
			// In
			resource *metadata.Resource
			href     string

			// Out
			variables []*metadata.PathVariable
			err       error

			// Test data
			r1       = regexp.MustCompile(`/clouds`)
			r2       = regexp.MustCompile(`/clouds/([^/]+)`)
			r3       = regexp.MustCompile(`/clouds/([^/]+)/instances/([^/]+)`)
			pattern1 = metadata.PathPattern{"GET", "", []string{}, r1}
			pattern2 = metadata.PathPattern{"GET", "%s", []string{"a"}, r2}
			pattern3 = metadata.PathPattern{"GET", "%s%s", []string{"a", "b"}, r3}
			a        = metadata.Action{PathPatterns: []*metadata.PathPattern{&pattern1}}
			b        = metadata.Action{PathPatterns: []*metadata.PathPattern{&pattern2}}
			c        = metadata.Action{PathPatterns: []*metadata.PathPattern{&pattern2, &pattern3}}
		)

		JustBeforeEach(func() {
			variables, err = resource.ExtractVariables(href)
		})

		Context("with a pattern that does not match", func() {
			BeforeEach(func() {
				resource = &metadata.Resource{Actions: []*metadata.Action{&a}}
				href = "/crouds"
			})

			It("returns an error", func() {
				Ω(variables).Should(BeEmpty())
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("with a pattern with no variable", func() {
			BeforeEach(func() {
				resource = &metadata.Resource{Actions: []*metadata.Action{&a}}
				href = "/clouds"
			})

			It("returns an empty array", func() {
				Ω(variables).Should(BeEmpty())
				Ω(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with a pattern that matches", func() {
			BeforeEach(func() {
				resource = &metadata.Resource{Actions: []*metadata.Action{&b}}
				href = "/clouds/foo"
			})

			It("returns the corresponding variable", func() {
				Ω(variables).Should(HaveLen(1))
				Ω(variables).Should(ContainElement(&metadata.PathVariable{"a", "foo"}))
				Ω(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with multiple patterns that matches", func() {
			BeforeEach(func() {
				resource = &metadata.Resource{Actions: []*metadata.Action{&b, &c}}
				href = "/clouds/foo/instances/bar"
			})

			It("returns the variables for the longest path", func() {
				Ω(variables).Should(HaveLen(2))
				Ω(variables).Should(ContainElement(&metadata.PathVariable{"a", "foo"}))
				Ω(variables).Should(ContainElement(&metadata.PathVariable{"b", "bar"}))
				Ω(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
