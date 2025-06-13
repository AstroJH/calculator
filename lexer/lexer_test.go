package lexer

import(
	"testing"

	"gosty/token"
)

func TestNextToken(t *testing.T)  {
	input := `=+(){},;`

	tests := []token.Token{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
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