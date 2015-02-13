package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ActionParam.Flags()", func() {
	var (
		actionParam *ActionParam
		flags       FlagDefs
		simpleParam *ActionParam
		otherParam  *ActionParam
	)

	BeforeEach(func() {
		s := BasicDataType("string")
		i := BasicDataType("int")
		simpleParam = &ActionParam{Name: "basic", Description: "description", Type: &s}
		otherParam = &ActionParam{Name: "other", Description: "other description", Type: &i}
	})

	JustBeforeEach(func() {
		flags = actionParam.Flags()
	})

	Context("for a basic type", func() {
		BeforeEach(func() {
			actionParam = simpleParam
		})
		It("returns a single flag", func() {
			f := FlagDefs{&FlagDef{
				simpleParam, "basic", "basic", "string", "description", false, false, false}}
			Ω(flags).Should(BeEquivalentTo(f))
		})
	})

	Context("for an array type", func() {
		BeforeEach(func() {
			arrayType := ArrayDataType{ElemType: simpleParam}
			actionParam = &ActionParam{Name: "array", Description: "array description", Type: &arrayType}
		})

		It("returns a single flag", func() {
			f := FlagDefs{&FlagDef{
				actionParam, "array", `array\.(\d+)`, "string", "array description", true, false, false}}
			Ω(flags).Should(BeEquivalentTo(f))
		})
	})

	Context("for a simple object type", func() {
		BeforeEach(func() {
			objectType := ObjectDataType{
				Name:   "ObjectType",
				Fields: []*ActionParam{simpleParam},
			}
			actionParam = &ActionParam{Name: "object", Type: &objectType}
		})

		It("returns a single flag", func() {
			f := FlagDefs{&FlagDef{
				simpleParam, "object.basic", `object.basic`, "string", "description", false, false, false}}
			Ω(flags).Should(BeEquivalentTo(f))
		})
	})

	Context("for a enum type", func() {
		BeforeEach(func() {
			enumType := new(EnumerableDataType)
			actionParam = &ActionParam{Name: "enum", Type: enumType, Description: "enum description"}
		})

		It("returns a single flag", func() {
			f := FlagDefs{&FlagDef{
				actionParam, "enum", `enum\.([a-z0-9_]+)`, "string", "enum description", false, true, false}}
			Ω(flags).Should(BeEquivalentTo(f))
		})
	})

	Context("for a one level deep object type", func() {

		BeforeEach(func() {
			objectType := ObjectDataType{
				Name:   "ObjectType",
				Fields: []*ActionParam{simpleParam, otherParam},
			}
			actionParam = &ActionParam{Name: "object", Type: &objectType}
		})

		It("returns two flags", func() {
			f := FlagDefs{
				&FlagDef{simpleParam, "object.basic", `object.basic`, "string", "description", false, false, false},
				&FlagDef{otherParam, "object.other", `object.other`, "int", "other description", false, false, false},
			}
			Ω(flags).Should(BeEquivalentTo(f))
		})
	})

	Context("for a recursive object", func() {

		var simpleParam2 *ActionParam

		BeforeEach(func() {
			childType1 := ObjectDataType{
				Name:   "ChildType1",
				Fields: []*ActionParam{simpleParam, otherParam},
			}
			child1 := &ActionParam{Name: "child1", Description: "child1 desc", Type: &childType1}
			st := BasicDataType("string")
			simpleParam2 = &ActionParam{Name: "simple2", Description: "simple2 desc", Type: &st}
			childType2 := ObjectDataType{
				Name:   "ChildType2",
				Fields: []*ActionParam{simpleParam2},
			}
			child2 := &ActionParam{Name: "child2", Description: "child2 desc", Type: &childType2}
			parentType := ObjectDataType{
				Name:   "ParentType",
				Fields: []*ActionParam{child1, child2},
			}
			actionParam = &ActionParam{Name: "parent", Type: &parentType}
		})

		It("returns three flags", func() {
			f := FlagDefs{
				&FlagDef{simpleParam, "parent.child1.basic", `parent.child1.basic`, "string", "description", false, false, false},
				&FlagDef{otherParam, "parent.child1.other", `parent.child1.other`, "int", "other description", false, false, false},
				&FlagDef{simpleParam2, "parent.child2.simple2", `parent.child2.simple2`, "string", "simple2 desc", false, false, false},
			}
			Ω(flags).Should(BeEquivalentTo(f))
		})
	})

	Context("for a heterogeneous object", func() {

		var enumParam, simpleParam2, grandchildParam *ActionParam

		BeforeEach(func() {
			enumParam = &ActionParam{Name: "enum", Description: "enum desc", Type: new(EnumerableDataType)}
			childType1 := ObjectDataType{Name: "ChildType1", Fields: []*ActionParam{simpleParam, enumParam}}
			childParam1 := &ActionParam{Name: "child1", Description: "child1 desc", Type: &childType1}

			st := BasicDataType("string")
			simpleParam2 = &ActionParam{Name: "simple2", Description: "simple2 desc", Type: &st}
			grandchildParam = &ActionParam{Name: "grandchildfield", Description: "grandchildfield", Type: &st}
			grandChild := &ActionParam{Name: "grandchild", Description: "grandchild desc", Type: &ObjectDataType{Name: "GrandChild", Fields: []*ActionParam{grandchildParam}}}
			array := &ActionParam{Name: "array", Description: "array desc", Type: &ArrayDataType{ElemType: grandChild}}
			childType2 := ObjectDataType{Name: "ChildType2", Fields: []*ActionParam{simpleParam2, array}}
			child2 := &ActionParam{Name: "child2", Description: "child2 desc", Type: &childType2}

			parentType := ObjectDataType{
				Name:   "ParentType",
				Fields: []*ActionParam{childParam1, child2},
			}
			actionParam = &ActionParam{Name: "parent", Type: &parentType}
		})

		It("returns all flags", func() {
			f := FlagDefs{
				&FlagDef{simpleParam, "parent.child1.basic", `parent.child1.basic`, "string", "description", false, false, false},
				&FlagDef{enumParam, "parent.child1.enum", `parent\.child1\.enum\.([a-z0-9_]+)`, "string", "enum desc", false, true, false},
				&FlagDef{simpleParam2, "parent.child2.simple2", `parent.child2.simple2`, "string", "simple2 desc", false, false, false},
				&FlagDef{grandchildParam, "parent.child2.grandchild.grandchildfield", `parent\.child2\.grandchild\.(\d+)\.grandchildfield`, "string", "grandchildfield", true, false, false},
			}
			Ω(flags).Should(BeEquivalentTo(f))
		})
	})

})
