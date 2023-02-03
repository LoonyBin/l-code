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
	ID       *string `(@Ident)?`
	Template *string `("<" "<" @Ident)?`
}

// Value represents a value in the EBNF syntax.
type Value struct {
	String *string   `  @String`
	Float  *float64  `| @Float`
	Int    *int      `| @Int`
	Object *ValueObj `| @@`
}

// ValueObj represents an object value in the EBNF syntax.
type ValueObj struct {
	Statements []*Statement `"{" @@* "}"`
}

var (
	Parser *participle.Parser[Program]
)

func init() {
	Parser, _ = participle.Build[Program]()
}

func ParseString(input string) (*Program, error) {
	return Parser.ParseString("", input)
}
