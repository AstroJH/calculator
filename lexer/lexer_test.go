package lexer

import (
	"testing"

	"calculator/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []token.Token{
		{
			Type:    token.ASSIGN,
			Literal: "=",
		},
		{
			Type:    token.PLUS,
			Literal: "+",
		},
		{
			Type:    token.LPAREN,
			Literal: "(",
		},
		{
			Type:    token.RPAREN,
			Literal: ")",
		},
		{
			Type:    token.LBRACE,
			Literal: "{",
		},
		{
			Type:    token.RBRACE,
			Literal: "}",
		},
		{
			Type:    token.COMMA,
			Literal: ",",
		},
		{
			Type:    token.SEMICOLON,
			Literal: ";",
		},
		{
			Type:    token.EOF,
			Literal: "",
		},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		for tok.Type != tt.Type {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q",
				i, tt.Type, tok.Type)
		}

		if tok.Literal != tt.Literal {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.Literal, tok.Literal)
		}
	}
}
