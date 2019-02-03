package lexer

import "interpreter-in-go/token"

type Lexer struct {
	input        string
	position     int // 入力における現在の位置
	readPosition int // これから読み込む位置（現在の文字の次)
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// 次の１文字をよんでinput文字列の現在位置を進める
// ASCII文字だけに対応 unicode全体をカバーしていない
func (l *Lexer) readChar() {
	// 入力が終端に達したかどうか 
	if l.readPosition >= len(l.input) {
		// 終端に達した場合 ０
		// 0はASCIIコードのNUL文字に対応
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.ASSIGN, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		// 終端として扱う
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type: tokenType,
		Literal: string(ch)
	}
}