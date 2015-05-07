package rsapi_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rightscale/rsc/rsapi"
)

var _ = Describe("normalize", func() {
	var payload rsapi.ApiParams
	var name string
	var value interface{}

	var res rsapi.ApiParams
	var resErr error

	JustBeforeEach(func() {
		res, resErr = rsapi.Normalize(payload, name, value)
	})

	BeforeEach(func() {
		payload = rsapi.ApiParams{}
	})

	Describe("with a simple string", func() {
		BeforeEach(func() {
			name = "val"
			value = "foo"
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			Ω(res).Should(Equal(rsapi.ApiParams{"val": "foo"}))
		})
	})

	Describe("with an int", func() {
		BeforeEach(func() {
			name = "val"
			value = 42
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			Ω(res).Should(Equal(rsapi.ApiParams{"val": 42}))
		})
	})

	Describe("with a float", func() {
		BeforeEach(func() {
			name = "val"
			value = 42.5
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			Ω(res).Should(Equal(rsapi.ApiParams{"val": 42.5}))
		})
	})

	Describe("with a bool", func() {
		BeforeEach(func() {
			name = "val"
			value = true
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			Ω(res).Should(Equal(rsapi.ApiParams{"val": true}))
		})
	})

	Describe("with a simple array", func() {
		BeforeEach(func() {
			name = "val[]"
			value = []string{"foo"}
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			Ω(res).Should(Equal(rsapi.ApiParams{"val": []string{"foo"}}))
		})
	})

	Describe("with a simple map", func() {
		BeforeEach(func() {
			name = "val[a]"
			value = "foo"
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			Ω(res).Should(Equal(rsapi.ApiParams{"val": rsapi.ApiParams{"a": "foo"}}))
		})
	})

	Describe("with a map of arrays", func() {
		BeforeEach(func() {
			name = "val[a][]"
			value = []interface{}{"foo"}
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			expected := rsapi.ApiParams{
				"val": rsapi.ApiParams{
					"a": []interface{}{"foo"},
				},
			}
			Ω(res).Should(Equal(expected))
		})
	})

	Describe("with a map of arrays with existing values", func() {
		BeforeEach(func() {
			name = "val[a][]"
			value = []string{"foo", "bar"}
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			expected := rsapi.ApiParams{
				"val": rsapi.ApiParams{
					"a": []string{"foo", "bar"},
				},
			}
			Ω(res).Should(Equal(expected))
		})
	})

	Describe("with an array of maps", func() {
		BeforeEach(func() {
			name = "val[][a]"
			value = "foo"
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			expected := rsapi.ApiParams{
				"val": []interface{}{rsapi.ApiParams{"a": "foo"}},
			}
			Ω(res).Should(Equal(expected))
		})
	})

	Describe("with an array of maps of arrays", func() {
		BeforeEach(func() {
			name = "val[][a][]"
			value = []interface{}{"foo"}
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			expected := rsapi.ApiParams{
				"val": []interface{}{rsapi.ApiParams{"a": []interface{}{"foo"}}},
			}
			Ω(res).Should(Equal(expected))
		})
	})

	Describe("with an array of maps with existing keys", func() {
		BeforeEach(func() {
			name = "val[][b]"
			value = "bar"
			payload = rsapi.ApiParams{
				"val": []interface{}{rsapi.ApiParams{"a": "foo"}},
			}
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			expected := rsapi.ApiParams{
				"val": []interface{}{rsapi.ApiParams{"a": "foo", "b": "bar"}},
			}
			Ω(res).Should(Equal(expected))
		})
	})

	Describe("with an array of maps with existing keys with more than one element", func() {
		BeforeEach(func() {
			name = "val[][b]"
			value = "baz"
			payload = rsapi.ApiParams{
				"val": []interface{}{rsapi.ApiParams{"a": "foo", "b": "bar"}},
			}
		})
		It("sets the value", func() {
			Ω(resErr).ShouldNot(HaveOccurred())
			expected := rsapi.ApiParams{
				"val": []interface{}{
					rsapi.ApiParams{"a": "foo", "b": "bar"},
					rsapi.ApiParams{"b": "baz"},
				},
			}
			Ω(res).Should(Equal(expected))
		})
	})
})
