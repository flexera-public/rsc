package cmds_test

import (
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rightscale/rsc/cmds"
)

var _ = Describe("ActionCmd", func() {

	Context("PathPattern Substitute", func() {
		var (
			// In
			pattern   *cmds.PathPattern
			variables []*cmds.PathVariable

			// Out
			path  string
			names []string

			// Test data
			a = cmds.PathVariable{"a", "1"}
			b = cmds.PathVariable{"b", "2"}
			c = cmds.PathVariable{"c", "3"}
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
				variables = []*cmds.PathVariable{&a, &b, &c}
				pattern = &cmds.PathPattern{p, []string{}, r}
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
				variables = []*cmds.PathVariable{&a, &b, &c}
				pattern = &cmds.PathPattern{p, []string{"a"}, r}
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
				variables = []*cmds.PathVariable{&a, &b, &c}
				pattern = &cmds.PathPattern{p, []string{"a", "d"}, r}
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
			action    *cmds.ActionCmd
			variables []*cmds.PathVariable

			// Out
			url string
			err error

			// Test data
			prefix = "/a/path/pattern/with/one/"
			p1     = &cmds.PathPattern{prefix + "%s", []string{"a"}, nil}
			p2     = &cmds.PathPattern{"%s%s", []string{"a", "b"}, nil}
			a      = cmds.PathVariable{"a", "1"}
			b      = cmds.PathVariable{"b", "2"}
		)

		JustBeforeEach(func() {
			url, err = action.Url(variables)
		})

		Context("with a single matching pattern", func() {
			BeforeEach(func() {
				action = &cmds.ActionCmd{
					PathPatterns: []*cmds.PathPattern{p1},
				}
				variables = []*cmds.PathVariable{&a}
			})

			It("returns the URL", func() {
				Ω(url).Should(Equal(prefix + a.Value))
				Ω(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with a multiple matching pattern", func() {
			BeforeEach(func() {
				action = &cmds.ActionCmd{
					PathPatterns: []*cmds.PathPattern{p1, p2},
				}
				variables = []*cmds.PathVariable{&a, &b}
			})

			It("returns the URL that matches the most variables", func() {
				Ω(url).Should(Equal(a.Value + b.Value))
				Ω(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with the reverse matching pattern", func() {
			BeforeEach(func() {
				action = &cmds.ActionCmd{
					PathPatterns: []*cmds.PathPattern{p2, p1},
				}
				variables = []*cmds.PathVariable{&a, &b}
			})

			It("returns the URL that matches the most variables", func() {
				Ω(url).Should(Equal(a.Value + b.Value))
				Ω(err).ShouldNot(HaveOccurred())
			})
		})

		Context("with no matching pattern", func() {
			BeforeEach(func() {
				action = &cmds.ActionCmd{
					PathPatterns: []*cmds.PathPattern{p2},
				}
				variables = []*cmds.PathVariable{&a}
			})

			It("returns an error", func() {
				Ω(url).Should(BeEmpty())
				Ω(err).Should(HaveOccurred())
			})
		})
	})
})
