package parser_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "l-code/parser"
)

var _ = Describe("Parser", func() {
	Context("With Key only", func() {
		var ast *Program

		BeforeEach(func() {
			input := `key`
			var err error
			ast, err = Parser.ParseString("", input)
			Expect(err).To(BeNil())
		})

		It("Should parse the key", func() {
			Expect(ast.Statements[0].Key).To(Equal("key"))
		})
	})

	Context("With String Value", func() {
		var ast *Program

		BeforeEach(func() {
			input := `key "value"`
			var err error
			ast, err = Parser.ParseString("", input)
			Expect(err).To(BeNil())
		})

		It("Should parse the key", func() {
			Expect(ast.Statements[0].Key).To(Equal("key"))
		})

		It("Should parse the value", func() {
			Expect(*ast.Statements[0].Values[0].String).To(Equal("\"value\""))
		})
	})

	Context("With Integer Value", func() {
		var ast *Program

		BeforeEach(func() {
			input := `key 4`
			var err error
			ast, err = Parser.ParseString("", input)
			Expect(err).To(BeNil())
		})

		It("Should parse the key", func() {
			Expect(ast.Statements[0].Key).To(Equal("key"))
		})

		It("Should parse the value", func() {
			Expect(*ast.Statements[0].Values[0].Int).To(Equal(4))
		})
	})

	Context("With Float Value", func() {
		var ast *Program

		BeforeEach(func() {
			input := `key 4.5`
			var err error
			ast, err = Parser.ParseString("", input)
			Expect(err).To(BeNil())
		})

		It("Should parse the key", func() {
			Expect(ast.Statements[0].Key).To(Equal("key"))
		})

		It("Should parse the value", func() {
			Expect(*ast.Statements[0].Values[0].Float).To(Equal(4.5))
		})
	})

	Context("With Object Value", func() {
		var ast *Program

		BeforeEach(func() {
			input := `key { object "value" }`
			var err error
			ast, err = Parser.ParseString("", input)
			Expect(err).To(BeNil())
		})

		It("Should parse the key", func() {
			Expect(ast.Statements[0].Key).To(Equal("key"))
		})

		It("Should parse the value", func() {
			object := ast.Statements[0].Values[0].Object
			Expect(object.Statements[0].Key).To(Equal("object"))
			Expect(*object.Statements[0].Values[0].String).To(Equal("\"value\""))
		})
	})

	Context("With Multiline Object Value", func() {
		var ast *Program

		BeforeEach(func() {
			input := `key {
				object "value"
				foo "bar"
			}`
			var err error
			ast, err = Parser.ParseString("", input)
			Expect(err).To(BeNil())
		})

		It("Should parse the key", func() {
			Expect(ast.Statements[0].Key).To(Equal("key"))
		})

		It("Should parse the first line", func() {
			object := ast.Statements[0].Values[0].Object
			Expect(object.Statements[0].Key).To(Equal("object"))
			Expect(*object.Statements[0].Values[0].String).To(Equal("\"value\""))
		})

		It("Should parse the second line", func() {
			object := ast.Statements[0].Values[0].Object
			Expect(object.Statements[1].Key).To(Equal("foo"))
			Expect(*object.Statements[1].Values[0].String).To(Equal("\"bar\""))
		})
	})

	Context("With multiple values", func() {
		var ast *Program

		BeforeEach(func() {
			input := `key "string", 4.5`
			var err error
			ast, err = Parser.ParseString("", input)
			Expect(err).To(BeNil())
		})

		It("Should parse the key", func() {
			Expect(ast.Statements[0].Key).To(Equal("key"))
		})

		It("Should parse the first value", func() {
			Expect(*ast.Statements[0].Values[0].String).To(Equal(`"string"`))
		})

		It("Should parse the second value", func() {
			Expect(*ast.Statements[0].Values[1].Float).To(Equal(4.5))
		})
	})

	Context("With ID", func() {
		var ast *Program

		BeforeEach(func() {
			input := `key id "string"`
			var err error
			ast, err = Parser.ParseString("", input)
			Expect(err).To(BeNil())
		})

		It("Should parse the key", func() {
			Expect(ast.Statements[0].Key).To(Equal("key"))
		})

		It("Should parse the id", func() {
			Expect(*ast.Statements[0].Metadata.ID).To(Equal("id"))
		})

		It("Should parse the value", func() {
			Expect(*ast.Statements[0].Values[0].String).To(Equal(`"string"`))
		})
	})

	Context("With ID and multiple values", func() {
		var ast *Program

		BeforeEach(func() {
			input := `key id "string", 4.5`
			var err error
			ast, err = Parser.ParseString("", input)
			Expect(err).To(BeNil())
		})

		It("Should parse the key", func() {
			Expect(ast.Statements[0].Key).To(Equal("key"))
		})

		It("Should parse the id", func() {
			Expect(*ast.Statements[0].Metadata.ID).To(Equal("id"))
		})

		It("Should parse the first value", func() {
			Expect(*ast.Statements[0].Values[0].String).To(Equal(`"string"`))
		})

		It("Should parse the second value", func() {
			Expect(*ast.Statements[0].Values[1].Float).To(Equal(4.5))
		})
	})

	Context("With template", func() {
		var ast *Program

		BeforeEach(func() {
			input := `key <<id "string", 4.5`
			var err error
			ast, err = Parser.ParseString("", input)
			Expect(err).To(BeNil())
		})

		It("Should parse the key", func() {
			Expect(ast.Statements[0].Key).To(Equal("key"))
		})

		It("Should parse the template id", func() {
			Expect(*ast.Statements[0].Metadata.Template).To(Equal("id"))
		})

		It("Should parse the first value", func() {
			Expect(*ast.Statements[0].Values[0].String).To(Equal(`"string"`))
		})

		It("Should parse the second value", func() {
			Expect(*ast.Statements[0].Values[1].Float).To(Equal(4.5))
		})
	})

	Context("With ID and template", func() {
		var ast *Program

		BeforeEach(func() {
			input := `key id << template_id "string", 4.5`
			var err error
			ast, err = Parser.ParseString("", input)
			Expect(err).To(BeNil())
		})

		It("Should parse the key", func() {
			Expect(ast.Statements[0].Key).To(Equal("key"))
		})

		It("Should parse the id", func() {
			Expect(*ast.Statements[0].Metadata.ID).To(Equal("id"))
		})

		It("Should parse the template id", func() {
			Expect(*ast.Statements[0].Metadata.Template).To(Equal("template_id"))
		})

		It("Should parse the first value", func() {
			Expect(*ast.Statements[0].Values[0].String).To(Equal(`"string"`))
		})

		It("Should parse the second value", func() {
			Expect(*ast.Statements[0].Values[1].Float).To(Equal(4.5))
		})
	})

	XContext("With Multiple Lines", func() {
		var ast *Program

		BeforeEach(func() {
			input := `
				set
				key "value"
				foo 15
				bar 12.5
				object {}
			`
			var err error
			ast, err = Parser.ParseString("", input)
			Expect(err).To(BeNil())
		})

		It("Should parse the key", func() {
			Expect(ast.Statements[0].Key).To(Equal("key"))
		})

		It("Should parse the first line", func() {
			object := ast.Statements[0].Values[0].Object
			Expect(object.Statements[0].Key).To(Equal("object"))
			Expect(*object.Statements[0].Values[0].String).To(Equal("\"value\""))
		})

		It("Should parse the second line", func() {
			object := ast.Statements[0].Values[0].Object
			Expect(object.Statements[1].Key).To(Equal("foo"))
			Expect(*object.Statements[1].Values[0].String).To(Equal("\"bar\""))
		})
	})

})
