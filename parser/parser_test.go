package parser_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "l-code/parser"
)

var _ = Describe("Parser", func() {
	Context("Given input string 'key", func() {
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

	Context("Given input string 'key \"value\"'", func() {
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

	Context("Given input string 'key 4'", func() {
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

	Context("Given input string 'key 4.5'", func() {
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

	Context("Given input string 'key \"string\", 4.5'", func() {
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

	Context("Given input string 'key \"id\" \"string\"", func() {
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
			Expect(ast.Statements[0].Metadata.ID()).To(Equal("id"))
		})

		It("Should parse the value", func() {
			Expect(*ast.Statements[0].Values[0].String).To(Equal(`"string"`))
		})
	})

	Context("Given input string 'key id \"string\", 4.5'", func() {
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
			Expect(ast.Statements[0].Metadata.ID()).To(Equal("id"))
		})

		It("Should parse the first value", func() {
			Expect(*ast.Statements[0].Values[0].String).To(Equal(`"string"`))
		})

		It("Should parse the second value", func() {
			Expect(*ast.Statements[0].Values[1].Float).To(Equal(4.5))
		})
	})

	Context("Given input string 'key template: id \"string\", 4.5'", func() {
		var ast *Program

		BeforeEach(func() {
			input := `key template: id "string", 4.5`
			var err error
			ast, err = Parser.ParseString("", input)
			Expect(err).To(BeNil())
		})

		It("Should parse the key", func() {
			Expect(ast.Statements[0].Key).To(Equal("key"))
		})

		It("Should parse the template id", func() {
			Expect(ast.Statements[0].Metadata.Template()).To(Equal("id"))
		})

		It("Should parse the first value", func() {
			Expect(*ast.Statements[0].Values[0].String).To(Equal(`"string"`))
		})

		It("Should parse the second value", func() {
			Expect(*ast.Statements[0].Values[1].Float).To(Equal(4.5))
		})
	})

	Context("Given input string 'key id, template: template_id \"string\", 4.5'", func() {
		var ast *Program

		BeforeEach(func() {
			input := `key id, template: template_id "string", 4.5`
			var err error
			ast, err = Parser.ParseString("", input)
			Expect(err).To(BeNil())
		})

		It("Should parse the key", func() {
			Expect(ast.Statements[0].Key).To(Equal("key"))
		})

		It("Should parse the id", func() {
			Expect(ast.Statements[0].Metadata.ID()).To(Equal("id"))
		})

		It("Should parse the template id", func() {
			Expect(ast.Statements[0].Metadata.Template()).To(Equal("template_id"))
		})

		It("Should parse the first value", func() {
			Expect(*ast.Statements[0].Values[0].String).To(Equal(`"string"`))
		})

		It("Should parse the second value", func() {
			Expect(*ast.Statements[0].Values[1].Float).To(Equal(4.5))
		})
	})
})
