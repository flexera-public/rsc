package metadata_test

import (
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rightscale/rsc/metadata"
)

var _ = Describe("Action", func() {

	Context("PathPattern Substitute", func() {
		var (
			// In
			pattern   *metadata.PathPattern
			variables []*metadata.PathVariable

			// Out
			path  string
			names []string

			// Test data
			a = metadata.PathVariable{"a", "1"}
			b = metadata.PathVariable{"b", "2"}
			c = metadata.PathVariable{"c", "3"}
		)

		JustBeforeEach(func() {
			path, names = pattern.Substitute(variables)
		})

		Context("with a pattern with no variable", func() {
			var (
				p = "/a/path/pattern/with/no/variable"
				r = regexp.MustCompile(p)
			)

			BeforeEach(func() {
				variables = []*metadata.PathVariable{&a, &b, &c}
				pattern = &metadata.PathPattern{"GET", p, []string{}, r}
			})

			It("returns the pattern as is", func() {
				Ω(path).Should(Equal(p))
				Ω(names).Should(BeEmpty())
			})
		})

		Context("with a pattern that matches", func() {
			var (
				prefix = "/a/path/pattern/with/one/"
				p      = prefix + "%s"
				r      = regexp.MustCompile("prefix + [[:alnum:]]")
			)

			BeforeEach(func() {
				variables = []*metadata.PathVariable{&a, &b, &c}
				pattern = &metadata.PathPattern{"GET", p, []string{"a"}, r}
			})

			It("returns the substituted pattern with the matched variables", func() {
				Ω(path).Should(Equal(prefix + a.Value))
				Ω(names).Should(Equal([]string{"a"}))
			})
		})

		Context("with a pattern that does not match", func() {
			var (
				prefix = "/a/path/pattern/with/one/"
				p      = prefix + "%s"
				r      = regexp.MustCompile("prefix + [[:alnum:]]")
			)

			BeforeEach(func() {
				variables = []*metadata.PathVariable{&a, &b, &c}
				pattern = &metadata.PathPattern{"GET", p, []string{"a", "d"}, r}
			})

			It("returns an empty string and the unmatched variables", func() {
				Ω(path).Should(BeEmpty())
				Ω(names).Should(Equal([]string{"d"}))
			})
		})
	})

	Context("Action Url", func() {
		var (
			// In
			action    *metadata.Action
			variables []*metadata.PathVariable

			// Out
			url string
			err error

			// Test data
			prefix = "/a/path/pattern/with/one/"
			p1     = &metadata.PathPattern{"GET", prefix + "%s", []string{"a"}, nil}
			p2     = &metadata.PathPattern{"GET", "%s%s", []string{"a", "b"}, nil}
			a      = metadata.PathVariable{"a", "1"}
			b      = metadata.PathVariable{"b", "2"}
		)

		JustBeforeEach(func() {
			var p *metadata.ActionPath
			p, err = action.URL(variables)
			if err == nil {
				url = p.Path
			}
		})

		Context("with a single matching pattern", func() {
			BeforeEach(func() {
				action = &metadata.Action{
					PathPatterns: []*metadata.PathPattern{p1},
				}
				variables = []*metadata.PathVariable{&a}
			})

			It("returns the URL", func() {
				Ω(url).Should(Equal(prefix + a.Value))
				Ω(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with a multiple matching pattern", func() {
			BeforeEach(func() {
				action = &metadata.Action{
					PathPatterns: []*metadata.PathPattern{p1, p2},
				}
				variables = []*metadata.PathVariable{&a, &b}
			})

			It("returns the URL that matches the most variables", func() {
				Ω(url).Should(Equal(a.Value + b.Value))
				Ω(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with the reverse matching pattern", func() {
			BeforeEach(func() {
				action = &metadata.Action{
					PathPatterns: []*metadata.PathPattern{p2, p1},
				}
				variables = []*metadata.PathVariable{&a, &b}
			})

			It("returns the URL that matches the most variables", func() {
				Ω(url).Should(Equal(a.Value + b.Value))
				Ω(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with no matching pattern", func() {
			BeforeEach(func() {
				action = &metadata.Action{
					PathPatterns: []*metadata.PathPattern{p2},
				}
				variables = []*metadata.PathVariable{&a}
			})

			It("returns an error", func() {
				Ω(err).Should(HaveOccurred())
			})
		})
	})
})
