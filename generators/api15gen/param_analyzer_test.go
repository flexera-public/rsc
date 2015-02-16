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
		analyzer = NewAnalyzer(params)
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
			param := params[0]
			Ω(param.Name).Should(Equal("foo"))
			s := BasicDataType("string")
			Ω(param.Type).Should(BeEquivalentTo(&s))
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
			param := params[0]
			Ω(param.Name).Should(Equal("foo"))
			s := BasicDataType("string")
			item := ActionParam{
				Name:      "item",
				QueryName: "foo[item]",
				VarName:   "item",
				Type:      &s,
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
			param := params[0]
			Ω(param.Name).Should(Equal("foo"))
			s := BasicDataType("string")
			bar := ActionParam{
				Name:      "bar",
				QueryName: "foo[bar]",
				VarName:   "bar",
				Type:      &s,
			}
			i := BasicDataType("int")
			baz := ActionParam{
				Name:      "baz",
				QueryName: "foo[baz]",
				VarName:   "baz",
				Type:      &i,
			}
			Ω(param.Type).Should(BeEquivalentTo(
				&ObjectDataType{"Foo", []*ActionParam{&bar, &baz}}))
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
			param := params[0]
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
			param := params[0]
			Ω(param.Name).Should(Equal("foo"))
			s := BasicDataType("string")
			bar := ActionParam{
				Name:      "bar",
				QueryName: "foo[bar]",
				Type:      &s,
				VarName:   "bar",
			}
			baz := ActionParam{
				Name:      "baz",
				QueryName: "foo[baz]",
				Type:      new(EnumerableDataType),
				VarName:   "baz",
			}
			Ω(param.Type).Should(BeEquivalentTo(
				&ObjectDataType{"Foo", []*ActionParam{&bar, &baz}}))
		})
	})

	Context("with a hash star param", func() {
		BeforeEach(func() {
			path = ""
			params = map[string]interface{}{
				"foo":    map[string]interface{}{"class": "Hash"},
				"foo[*]": map[string]interface{}{"class": "String"},
			}
		})

		It("Analyze returns the parsed param", func() {
			analyzer.Analyze()
			params := analyzer.Params
			Ω(params).Should(HaveLen(1))
			param := params[0]
			Ω(param.Name).Should(Equal("foo"))
			Ω(param.Type).Should(BeEquivalentTo(new(EnumerableDataType)))
		})
	})

	Context("with a orphan star param", func() {
		BeforeEach(func() {
			path = ""
			params = map[string]interface{}{
				"foo[*]": map[string]interface{}{"class": "String"},
			}
		})

		It("Analyze returns the parsed param", func() {
			analyzer.Analyze()
			params := analyzer.Params
			Ω(params).Should(HaveLen(1))
			param := params[0]
			Ω(param.Name).Should(Equal("foo"))
			Ω(param.Type).Should(BeEquivalentTo(new(EnumerableDataType)))
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
			param := params[0]
			Ω(param.Name).Should(Equal("foo"))
			s := BasicDataType("string")
			bar := ActionParam{
				Name:      "bar",
				QueryName: "foo[][bar]",
				Type:      &s,
				VarName:   "bar",
			}
			i := BasicDataType("int")
			baz := ActionParam{
				Name:      "baz",
				QueryName: "foo[][baz]",
				Type:      &i,
				VarName:   "baz",
			}
			t := ObjectDataType{"Foo", []*ActionParam{&bar, &baz}}
			item := ActionParam{
				Name:      "item",
				QueryName: "foo[][item]",
				Type:      &t,
				VarName:   "item",
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
			param := params[0]
			Ω(param.Name).Should(Equal("foo"))
			s := BasicDataType("string")
			bar := ActionParam{
				Name:      "bar",
				QueryName: "foo[][bar]",
				Type:      &s,
				VarName:   "bar",
			}
			goo := ActionParam{
				Name:      "goo",
				QueryName: "foo[][baz][][goo]",
				Type:      &s,
				VarName:   "goo",
			}
			t := ObjectDataType{"FooBaz", []*ActionParam{&goo}}
			bazItem := ActionParam{
				Name:      "item",
				QueryName: "foo[][baz][][item]",
				Type:      &t,
				VarName:   "item",
			}
			baz := ActionParam{
				Name:      "baz",
				QueryName: "foo[][baz]",
				Type:      &ArrayDataType{&bazItem},
				VarName:   "baz",
			}
			p := ObjectDataType{"Foo", []*ActionParam{&bar, &baz}}
			item := ActionParam{
				Name:      "item",
				QueryName: "foo[][item]",
				Type:      &p,
				VarName:   "item",
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

			param := params[0]
			Ω(param.Name).Should(Equal("foo"))
			item := ActionParam{
				Name:      "item",
				QueryName: "foo[item]",
				Type:      &s,
				VarName:   "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&ArrayDataType{&item}))

			param = params[1]
			Ω(param.Name).Should(Equal("foo1"))
			Ω(param.Type).Should(BeEquivalentTo(new(EnumerableDataType)))

			param = params[2]
			Ω(param.Type).Should(BeEquivalentTo(&s))

			param = params[3]
			baz := ActionParam{
				Name:      "baz",
				QueryName: "foo3[baz]",
				Type:      &i,
				VarName:   "baz",
			}
			p := ObjectDataType{"Foo3", []*ActionParam{&baz}}
			Ω(param.Type).Should(BeEquivalentTo(&p))

			param = params[4]
			Ω(param.Name).Should(Equal("foo4"))
			bar := ActionParam{
				Name:      "bar",
				QueryName: "foo4[][bar]",
				Type:      &s,
				VarName:   "bar",
			}
			goo := ActionParam{
				Name:      "goo",
				QueryName: "foo4[][baz][][goo]",
				Type:      &s,
				VarName:   "goo",
			}
			zoo := ActionParam{
				Name:      "zoo",
				QueryName: "foo4[][baz][][zoo]",
				Type:      new(EnumerableDataType),
				VarName:   "zoo",
			}
			t := ObjectDataType{"Foo4Baz", []*ActionParam{&goo, &zoo}}
			bazItem := ActionParam{
				Name:      "item",
				QueryName: "foo4[][baz][][item]",
				Type:      &t,
				VarName:   "item",
			}
			baz = ActionParam{
				Name:      "baz",
				QueryName: "foo4[][baz]",
				Type:      &ArrayDataType{&bazItem},
				VarName:   "baz",
			}
			p = ObjectDataType{"Foo4", []*ActionParam{&bar, &baz}}
			item = ActionParam{
				Name:      "item",
				QueryName: "foo4[][item]",
				Type:      &p,
				VarName:   "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&ArrayDataType{&item}))
		})

	})

})
