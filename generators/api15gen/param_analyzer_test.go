package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/rightscale/rsc/generators/api15gen"
)

var _ = Describe("ParamAnalyzer", func() {
	var (
		path     string
		params   map[string]interface{}
		analyzer *main.ParamAnalyzer
	)

	JustBeforeEach(func() {
		analyzer = main.NewAnalyzer(params)
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
			s := main.BasicDataType("string")
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
			s := main.BasicDataType("string")
			item := main.ActionParam{
				Name:      "item",
				QueryName: "foo[item]",
				VarName:   "item",
				Type:      &s,
			}
			Ω(param.Type).Should(BeEquivalentTo(&main.ArrayDataType{&item}))
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
			s := main.BasicDataType("string")
			bar := main.ActionParam{
				Name:      "bar",
				QueryName: "foo[bar]",
				VarName:   "bar",
				Type:      &s,
			}
			i := main.BasicDataType("int")
			baz := main.ActionParam{
				Name:      "baz",
				QueryName: "foo[baz]",
				VarName:   "baz",
				Type:      &i,
			}
			Ω(param.Type).Should(BeEquivalentTo(
				&main.ObjectDataType{"Foo", []*main.ActionParam{&bar, &baz}}))
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
			Ω(param.Type).Should(BeEquivalentTo(new(main.EnumerableDataType)))
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
			s := main.BasicDataType("string")
			bar := main.ActionParam{
				Name:      "bar",
				QueryName: "foo[bar]",
				Type:      &s,
				VarName:   "bar",
			}
			baz := main.ActionParam{
				Name:      "baz",
				QueryName: "foo[baz]",
				Type:      new(main.EnumerableDataType),
				VarName:   "baz",
			}
			Ω(param.Type).Should(BeEquivalentTo(
				&main.ObjectDataType{"Foo", []*main.ActionParam{&bar, &baz}}))
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
			Ω(param.Type).Should(BeEquivalentTo(new(main.EnumerableDataType)))
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
			Ω(param.Type).Should(BeEquivalentTo(new(main.EnumerableDataType)))
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
			s := main.BasicDataType("string")
			bar := main.ActionParam{
				Name:      "bar",
				QueryName: "foo[][bar]",
				Type:      &s,
				VarName:   "bar",
			}
			i := main.BasicDataType("int")
			baz := main.ActionParam{
				Name:      "baz",
				QueryName: "foo[][baz]",
				Type:      &i,
				VarName:   "baz",
			}
			t := main.ObjectDataType{"Foo", []*main.ActionParam{&bar, &baz}}
			item := main.ActionParam{
				Name:      "item",
				QueryName: "foo[][item]",
				Type:      &t,
				VarName:   "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&main.ArrayDataType{&item}))
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
			s := main.BasicDataType("string")
			bar := main.ActionParam{
				Name:      "bar",
				QueryName: "foo[][bar]",
				Type:      &s,
				VarName:   "bar",
			}
			goo := main.ActionParam{
				Name:      "goo",
				QueryName: "foo[][baz][][goo]",
				Type:      &s,
				VarName:   "goo",
			}
			t := main.ObjectDataType{"Baz", []*main.ActionParam{&goo}}
			bazItem := main.ActionParam{
				Name:      "item",
				QueryName: "foo[][baz][][item]",
				Type:      &t,
				VarName:   "item",
			}
			baz := main.ActionParam{
				Name:      "baz",
				QueryName: "foo[][baz][]",
				Type:      &main.ArrayDataType{&bazItem},
				VarName:   "baz",
			}
			p := main.ObjectDataType{"Foo", []*main.ActionParam{&bar, &baz}}
			item := main.ActionParam{
				Name:      "item",
				QueryName: "foo[][item]",
				Type:      &p,
				VarName:   "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&main.ArrayDataType{&item}))
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

			s := main.BasicDataType("string")
			i := main.BasicDataType("int")

			param := params[0]
			Ω(param.Name).Should(Equal("foo"))
			item := main.ActionParam{
				Name:      "item",
				QueryName: "foo[item]",
				Type:      &s,
				VarName:   "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&main.ArrayDataType{&item}))

			param = params[1]
			Ω(param.Name).Should(Equal("foo1"))
			Ω(param.Type).Should(BeEquivalentTo(new(main.EnumerableDataType)))

			param = params[2]
			Ω(param.Type).Should(BeEquivalentTo(&s))

			param = params[3]
			baz := main.ActionParam{
				Name:      "baz",
				QueryName: "foo3[baz]",
				Type:      &i,
				VarName:   "baz",
			}
			p := main.ObjectDataType{"Foo3", []*main.ActionParam{&baz}}
			Ω(param.Type).Should(BeEquivalentTo(&p))

			param = params[4]
			Ω(param.Name).Should(Equal("foo4"))
			bar := main.ActionParam{
				Name:      "bar",
				QueryName: "foo4[][bar]",
				Type:      &s,
				VarName:   "bar",
			}
			goo := main.ActionParam{
				Name:      "goo",
				QueryName: "foo4[][baz][][goo]",
				Type:      &s,
				VarName:   "goo",
			}
			zoo := main.ActionParam{
				Name:      "zoo",
				QueryName: "foo4[][baz][][zoo]",
				Type:      new(main.EnumerableDataType),
				VarName:   "zoo",
			}
			t := main.ObjectDataType{"Baz", []*main.ActionParam{&goo, &zoo}}
			bazItem := main.ActionParam{
				Name:      "item",
				QueryName: "foo4[][baz][][item]",
				Type:      &t,
				VarName:   "item",
			}
			baz = main.ActionParam{
				Name:      "baz",
				QueryName: "foo4[][baz][]",
				Type:      &main.ArrayDataType{&bazItem},
				VarName:   "baz",
			}
			p = main.ObjectDataType{"Foo4", []*main.ActionParam{&bar, &baz}}
			item = main.ActionParam{
				Name:      "item",
				QueryName: "foo4[][item]",
				Type:      &p,
				VarName:   "item",
			}
			Ω(param.Type).Should(BeEquivalentTo(&main.ArrayDataType{&item}))
		})

	})

})
