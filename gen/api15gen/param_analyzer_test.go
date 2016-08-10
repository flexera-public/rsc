package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/rightscale/rsc/gen"
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
			s := gen.BasicDataType("string")
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
			s := gen.BasicDataType("string")
			item := gen.ActionParam{
				Name:      "item",
				QueryName: "foo[item]",
				VarName:   "item",
				Type:      &s,
			}
			Ω(param.Type).Should(BeEquivalentTo(&gen.ArrayDataType{&item}))
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
			s := gen.BasicDataType("string")
			bar := gen.ActionParam{
				Name:      "bar",
				QueryName: "foo[bar]",
				VarName:   "bar",
				Type:      &s,
			}
			i := gen.BasicDataType("int")
			baz := gen.ActionParam{
				Name:      "baz",
				QueryName: "foo[baz]",
				VarName:   "baz",
				Type:      &i,
			}
			Ω(param.Type).Should(BeEquivalentTo(
				&gen.ObjectDataType{"Foo", []*gen.ActionParam{&bar, &baz}}))
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
			Ω(param.Type).Should(BeEquivalentTo(new(gen.EnumerableDataType)))
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
			s := gen.BasicDataType("string")
			bar := gen.ActionParam{
				Name:      "bar",
				QueryName: "foo[bar]",
				Type:      &s,
				VarName:   "bar",
			}
			baz := gen.ActionParam{
				Name:      "baz",
				QueryName: "foo[baz]",
				Type:      new(gen.EnumerableDataType),
				VarName:   "baz",
			}
			Ω(param.Type).Should(BeEquivalentTo(
				&gen.ObjectDataType{"Foo", []*gen.ActionParam{&bar, &baz}}))
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
			Ω(param.Type).Should(BeEquivalentTo(new(gen.EnumerableDataType)))
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
			Ω(param.Type).Should(BeEquivalentTo(new(gen.EnumerableDataType)))
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
			s := gen.BasicDataType("string")
			bar := gen.ActionParam{
				Name:      "bar",
				QueryName: "foo[][bar]",
				Type:      &s,
				VarName:   "bar",
			}
			i := gen.BasicDataType("int")
			baz := gen.ActionParam{
				Name:      "baz",
				QueryName: "foo[][baz]",
				Type:      &i,
				VarName:   "baz",
			}
			t := gen.ObjectDataType{"Foo", []*gen.ActionParam{&bar, &baz}}
			item := gen.ActionParam{
				Name:      "item",
				QueryName: "foo[][item]",
				Type:      &t,
				VarName:   "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&gen.ArrayDataType{&item}))
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
			s := gen.BasicDataType("string")
			bar := gen.ActionParam{
				Name:      "bar",
				QueryName: "foo[][bar]",
				Type:      &s,
				VarName:   "bar",
			}
			goo := gen.ActionParam{
				Name:      "goo",
				QueryName: "foo[][baz][][goo]",
				Type:      &s,
				VarName:   "goo",
			}
			t := gen.ObjectDataType{"Baz", []*gen.ActionParam{&goo}}
			bazItem := gen.ActionParam{
				Name:      "item",
				QueryName: "foo[][baz][][item]",
				Type:      &t,
				VarName:   "item",
			}
			baz := gen.ActionParam{
				Name:      "baz",
				QueryName: "foo[][baz][]",
				Type:      &gen.ArrayDataType{&bazItem},
				VarName:   "baz",
			}
			p := gen.ObjectDataType{"Foo", []*gen.ActionParam{&bar, &baz}}
			item := gen.ActionParam{
				Name:      "item",
				QueryName: "foo[][item]",
				Type:      &p,
				VarName:   "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&gen.ArrayDataType{&item}))
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

			s := gen.BasicDataType("string")
			i := gen.BasicDataType("int")

			param := params[0]
			Ω(param.Name).Should(Equal("foo"))
			item := gen.ActionParam{
				Name:      "item",
				QueryName: "foo[item]",
				Type:      &s,
				VarName:   "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&gen.ArrayDataType{&item}))

			param = params[1]
			Ω(param.Name).Should(Equal("foo1"))
			Ω(param.Type).Should(BeEquivalentTo(new(gen.EnumerableDataType)))

			param = params[2]
			Ω(param.Type).Should(BeEquivalentTo(&s))

			param = params[3]
			baz := gen.ActionParam{
				Name:      "baz",
				QueryName: "foo3[baz]",
				Type:      &i,
				VarName:   "baz",
			}
			p := gen.ObjectDataType{"Foo3", []*gen.ActionParam{&baz}}
			Ω(param.Type).Should(BeEquivalentTo(&p))

			param = params[4]
			Ω(param.Name).Should(Equal("foo4"))
			bar := gen.ActionParam{
				Name:      "bar",
				QueryName: "foo4[][bar]",
				Type:      &s,
				VarName:   "bar",
			}
			goo := gen.ActionParam{
				Name:      "goo",
				QueryName: "foo4[][baz][][goo]",
				Type:      &s,
				VarName:   "goo",
			}
			zoo := gen.ActionParam{
				Name:      "zoo",
				QueryName: "foo4[][baz][][zoo]",
				Type:      new(gen.EnumerableDataType),
				VarName:   "zoo",
			}
			t := gen.ObjectDataType{"Baz", []*gen.ActionParam{&goo, &zoo}}
			bazItem := gen.ActionParam{
				Name:      "item",
				QueryName: "foo4[][baz][][item]",
				Type:      &t,
				VarName:   "item",
			}
			baz = gen.ActionParam{
				Name:      "baz",
				QueryName: "foo4[][baz][]",
				Type:      &gen.ArrayDataType{&bazItem},
				VarName:   "baz",
			}
			p = gen.ObjectDataType{"Foo4", []*gen.ActionParam{&bar, &baz}}
			item = gen.ActionParam{
				Name:      "item",
				QueryName: "foo4[][item]",
				Type:      &p,
				VarName:   "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&gen.ArrayDataType{&item}))
		})

	})

})
