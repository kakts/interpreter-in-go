package token

// TokenTypeの定義一覧
const (
	ILLEGAL = "ILLEGAL" // tokenが未知の文字
	EOF     = "EOF"

	// 識別子 + リテラル
	IDENT = "IDENT" // add, foober, x, y...
	INT   = "INT"   // 123456

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keyword
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
