package parser

import "github.com/alecthomas/participle/v2"

// Program is the main structure for the EBNF syntax.
type Program struct {
	Statements []*Statement `@@+`
}

// Statement represents a statement in the EBNF syntax.
type Statement struct {
	Key      string    `@Ident`
	Metadata *Metadata `@@?`
	Values   []*Value  `(@@ ("," @@)*)?`
}

type Metadata struct {
	IdAndTemplate *IdAndTemplate `  @@`
	TemplateOnly  *string        `| "template" ":" @Ident`
	IdOnly        *string        `| @Ident`
}

type IdAndTemplate struct {
	ID       string `@Ident ","`
	Template string `"template" ":" @Ident`
}

func (m *Metadata) ID() string {
	if m.IdOnly != nil {
		return *m.IdOnly
	}
	if m.IdAndTemplate != nil {
		return m.IdAndTemplate.ID
	}
	return ""

}

func (m *Metadata) Template() string {
	if m.TemplateOnly != nil {
		return *m.TemplateOnly
	}
	if m.IdAndTemplate != nil {
		return m.IdAndTemplate.Template
	}
	return ""
}

// Value represents a value in the EBNF syntax.
type Value struct {
	String *string  `  @String`
	Float  *float64 `| @Float`
	Int    *int     `| @Int`
	// Object *ValueObj `| @@`
}

// // ValueObj represents an object value in the EBNF syntax.
// type ValueObj struct {
// 	Statement []*Statement `"{" { @@ } "}"`
// }

var (
	Parser *participle.Parser[Program]
)

func init() {
	Parser, _ = participle.Build[Program]()
}

func ParseString(input string) (*Program, error) {
	return Parser.ParseString("", input)
}
