package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParamAnalyzer", func() {
	var (
		path     string
		params   map[string]interface{}
		analyzer *ParamAnalyzer
	)

	JustBeforeEach(func() {
		analyzer = NewAnalyzer(path, params)
	})

	Context("with an empty path and a simple param", func() {
		BeforeEach(func() {
			path = ""
			params = map[string]interface{}{"foo": map[string]interface{}{"class": "String"}}
		})

		It("Analyze returns the parsed param", func() {
			analyzer.Analyze()
			params := analyzer.Params
			Ω(params).Should(HaveLen(1))
			Ω(params).Should(HaveKey("foo"))
			param := params["foo"]
			Ω(param.Name).Should(Equal("foo"))
			s := BasicDataType("string")
			Ω(param.Type).Should(BeEquivalentTo(&s))
		})
	})

	Context("with a path and a no param", func() {
		BeforeEach(func() {
			path = "/foo/:bar"
			params = map[string]interface{}{}
		})

		It("Analyze extracts the path params", func() {
			analyzer.Analyze()
			params := analyzer.Params
			Ω(params).Should(HaveLen(1))
			Ω(params).Should(HaveKey("bar"))
			param := params["bar"]
			Ω(param.Name).Should(Equal("bar"))
			s := BasicDataType("string")
			Ω(param.Type).Should(BeEquivalentTo(&s))
		})
	})

	Context("with a both path params and payload params", func() {
		BeforeEach(func() {
			path = "/foo/:bar/baz/:id"
			params = map[string]interface{}{
				"foo": map[string]interface{}{"class": "String"},
				"goo": map[string]interface{}{"class": "Integer"},
			}
		})

		It("Analyze extracts both the path and payload params", func() {
			analyzer.Analyze()
			params := analyzer.Params
			Ω(params).Should(HaveLen(4))
			Ω(params).Should(HaveKey("bar"))
			Ω(params).Should(HaveKey("id"))
			Ω(params).Should(HaveKey("goo"))
			Ω(params).Should(HaveKey("foo"))
			param := params["bar"]
			Ω(param.Name).Should(Equal("bar"))
			s := BasicDataType("string")
			Ω(param.Type).Should(BeEquivalentTo(&s))
			param = params["id"]
			Ω(param.Name).Should(Equal("id"))
			Ω(param.Type).Should(BeEquivalentTo(&s))
			param = params["foo"]
			Ω(param.Name).Should(Equal("foo"))
			Ω(param.Type).Should(BeEquivalentTo(&s))
			param = params["goo"]
			Ω(param.Name).Should(Equal("goo"))
			i := BasicDataType("int")
			Ω(param.Type).Should(BeEquivalentTo(&i))
		})
	})

	Context("with a simple array param", func() {
		BeforeEach(func() {
			path = ""
			params = map[string]interface{}{"foo": map[string]interface{}{"class": "Array"}}
		})

		It("Analyze returns the parsed param", func() {
			analyzer.Analyze()
			params := analyzer.Params
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

	Context("with a simple hash param", func() {
		BeforeEach(func() {
			path = ""
			params = map[string]interface{}{
				"foo":      map[string]interface{}{"class": "Hash"},
				"foo[bar]": map[string]interface{}{"class": "String"},
				"foo[baz]": map[string]interface{}{"class": "Integer"},
			}
		})

		It("Analyze returns the parsed param", func() {
			analyzer.Analyze()
			params := analyzer.Params
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
			Ω(param.Type).Should(BeEquivalentTo(
				&ObjectDataType{"FooParam", []*ActionParam{&bar, &baz}}))
		})
	})

	Context("with a simple enumerable param", func() {
		BeforeEach(func() {
			path = ""
			params = map[string]interface{}{
				"foo":    map[string]interface{}{"class": "Enumerable"},
				"foo[*]": map[string]interface{}{"class": "String"},
			}
		})

		It("Analyze returns the parsed param", func() {
			analyzer.Analyze()
			params := analyzer.Params
			Ω(params).Should(HaveLen(1))
			Ω(params).Should(HaveKey("foo"))
			param := params["foo"]
			Ω(param.Name).Should(Equal("foo"))
			Ω(param.Type).Should(BeEquivalentTo(new(EnumerableDataType)))
		})
	})

	Context("with a hash of enumerable params", func() {
		BeforeEach(func() {
			path = ""
			params = map[string]interface{}{
				"foo":         map[string]interface{}{"class": "Hash"},
				"foo[bar]":    map[string]interface{}{"class": "String"},
				"foo[baz]":    map[string]interface{}{"class": "Enumerable"},
				"foo[baz][*]": map[string]interface{}{"class": "String"},
			}
		})

		It("Analyze returns the parsed param", func() {
			analyzer.Analyze()
			params := analyzer.Params
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
			baz := ActionParam{
				Name:       "baz",
				Type:       new(EnumerableDataType),
				NativeName: "baz",
			}
			Ω(param.Type).Should(BeEquivalentTo(
				&ObjectDataType{"FooParam", []*ActionParam{&bar, &baz}}))
		})
	})

	Context("with an array of hashes", func() {
		BeforeEach(func() {
			path = ""
			params = map[string]interface{}{
				"foo":        map[string]interface{}{"class": "Array"},
				"foo[][bar]": map[string]interface{}{"class": "String"},
				"foo[][baz]": map[string]interface{}{"class": "Integer"},
			}
		})

		It("Analyze returns the parsed param", func() {
			analyzer.Analyze()
			params := analyzer.Params
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
			t := ObjectDataType{"FooItemParam", []*ActionParam{&bar, &baz}}
			item := ActionParam{
				Name:       "item",
				Type:       &t,
				NativeName: "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&ArrayDataType{&item}))
		})
	})

	Context("with an array of hashes with sub-array", func() {
		BeforeEach(func() {
			path = ""
			params = map[string]interface{}{
				"foo":               map[string]interface{}{"class": "Array"},
				"foo[][bar]":        map[string]interface{}{"class": "String"},
				"foo[][baz]":        map[string]interface{}{"class": "Array"},
				"foo[][baz][][goo]": map[string]interface{}{"class": "String"},
			}
		})

		It("Analyze returns the parsed param", func() {
			analyzer.Analyze()
			params := analyzer.Params
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
			goo := ActionParam{
				Name:       "goo",
				Type:       &s,
				NativeName: "goo",
			}
			t := ObjectDataType{"BazItemParam", []*ActionParam{&goo}}
			bazItem := ActionParam{
				Name:       "item",
				Type:       &t,
				NativeName: "item",
			}
			baz := ActionParam{
				Name:       "baz",
				Type:       &ArrayDataType{&bazItem},
				NativeName: "baz",
			}
			p := ObjectDataType{"FooItemParam", []*ActionParam{&bar, &baz}}
			item := ActionParam{
				Name:       "item",
				Type:       &p,
				NativeName: "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&ArrayDataType{&item}))
		})
	})

	Context("with a mix of arrays, enumerables and hashes", func() {
		BeforeEach(func() {
			path = ""
			params = map[string]interface{}{
				"foo4":                  map[string]interface{}{"class": "Array"},
				"foo4[][bar]":           map[string]interface{}{"class": "String"},
				"foo2":                  map[string]interface{}{"class": "String"},
				"foo4[][baz][][zoo][*]": map[string]interface{}{"class": "String"},
				"foo1[*]":               map[string]interface{}{"class": "String"},
				"foo4[][baz]":           map[string]interface{}{"class": "Array"},
				"foo3":                  map[string]interface{}{"class": "Hash"},
				"foo4[][baz][][goo]":    map[string]interface{}{"class": "String"},
				"foo3[baz]":             map[string]interface{}{"class": "Integer"},
				"foo1":                  map[string]interface{}{"class": "Enumerable"},
				"foo":                   map[string]interface{}{"class": "Array"},
				"foo4[][baz][][zoo]":    map[string]interface{}{"class": "Enumerable"},
			}
		})

		It("Analyze returns the parsed params", func() {
			analyzer.Analyze()
			params := analyzer.Params
			Ω(params).Should(HaveLen(5))

			s := BasicDataType("string")
			i := BasicDataType("int")

			Ω(params).Should(HaveKey("foo"))
			param := params["foo"]
			Ω(param.Name).Should(Equal("foo"))
			item := ActionParam{
				Name:       "item",
				Type:       &s,
				NativeName: "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&ArrayDataType{&item}))

			Ω(params).Should(HaveKey("foo1"))
			param = params["foo1"]
			Ω(param.Name).Should(Equal("foo1"))
			Ω(param.Type).Should(BeEquivalentTo(new(EnumerableDataType)))

			Ω(params).Should(HaveKey("foo2"))
			param = params["foo2"]
			Ω(param.Type).Should(BeEquivalentTo(&s))

			Ω(params).Should(HaveKey("foo3"))
			param = params["foo3"]
			baz := ActionParam{
				Name:       "baz",
				Type:       &i,
				NativeName: "baz",
			}
			p := ObjectDataType{"Foo3Param", []*ActionParam{&baz}}
			Ω(param.Type).Should(BeEquivalentTo(&p))

			Ω(params).Should(HaveKey("foo4"))
			param = params["foo4"]
			Ω(param.Name).Should(Equal("foo4"))
			bar := ActionParam{
				Name:       "bar",
				Type:       &s,
				NativeName: "bar",
			}
			goo := ActionParam{
				Name:       "goo",
				Type:       &s,
				NativeName: "goo",
			}
			zoo := ActionParam{
				Name:       "zoo",
				Type:       new(EnumerableDataType),
				NativeName: "zoo",
			}
			t := ObjectDataType{"BazItemParam", []*ActionParam{&goo, &zoo}}
			bazItem := ActionParam{
				Name:       "item",
				Type:       &t,
				NativeName: "item",
			}
			baz = ActionParam{
				Name:       "baz",
				Type:       &ArrayDataType{&bazItem},
				NativeName: "baz",
			}
			p = ObjectDataType{"Foo4ItemParam", []*ActionParam{&bar, &baz}}
			item = ActionParam{
				Name:       "item",
				Type:       &p,
				NativeName: "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&ArrayDataType{&item}))
		})

	})

})
