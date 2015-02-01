package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParamAnalyzer", func() {
	var (
		path   string
		params map[string]interface{}
	)

	Context("With an empty path and a simple param", func() {
		BeforeEach(func() {
			path = ""
			params = map[string]interface{}{"foo": map[string]interface{}{"class": "String"}}
		})

		It("should return the parsed param", func() {
			params := parseParams(path, params)
			Ω(params).Should(HaveLen(1))
			Ω(params).Should(HaveKey("foo"))
			param := params["foo"]
			Ω(param.Name).Should(Equal("foo"))
			s := BasicDataType("string")
			Ω(param.Type).Should(BeEquivalentTo(&s))
		})
	})

	Context("With a simple array param", func() {
		BeforeEach(func() {
			path = ""
			params = map[string]interface{}{"foo": map[string]interface{}{"class": "Array"}}
		})

		It("should return the parsed param", func() {
			params := parseParams(path, params)
			Ω(params).Should(HaveLen(1))
			Ω(params).Should(HaveKey("foo"))
			param := params["foo"]
			Ω(param.Name).Should(Equal("foo"))
			s := BasicDataType("string")
			item := ActionParam{
				Name:       "item",
				Type:       &s,
				NativeName: "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&ArrayDataType{&item}))
		})
	})

	Context("With an array of hashes", func() {
		BeforeEach(func() {
			path = ""
			params = map[string]interface{}{
				"foo":        map[string]interface{}{"class": "Array"},
				"foo[][bar]": map[string]interface{}{"class": "String"},
				"foo[][baz]": map[string]interface{}{"class": "Integer"},
			}
		})

		It("should return the parsed param", func() {
			params := parseParams(path, params)
			Ω(params).Should(HaveLen(1))
			Ω(params).Should(HaveKey("foo"))
			param := params["foo"]
			Ω(param.Name).Should(Equal("foo"))
			s := BasicDataType("string")
			bar := ActionParam{
				Name:       "bar",
				Type:       &s,
				NativeName: "bar",
			}
			i := BasicDataType("int")
			baz := ActionParam{
				Name:       "baz",
				Type:       &i,
				NativeName: "baz",
			}
			t := ObjectDataType{"fooItem", []*ActionParam{&bar, &baz}}
			item := ActionParam{
				Name:       "item",
				Type:       &t,
				NativeName: "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&ArrayDataType{&item}))
		})
	})
})
