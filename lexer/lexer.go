package lexer

import (
	"github.com/kakts/monkey/token"
)

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
	l.readPosition++
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	// 文字が続くまでよみすすめる
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 簡単のため整数のみ対応
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	// ホワイトスペース、タブ、改行以外の文字がでるまで読み進める
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// l.position, l.readPositionをインクリメンテせず、次の文字列を先読みする
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0;
	} else {
		// 現在のインデックスの次の文字を取得
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	// ホワイトスペースはスキップさせる
	l.skipWhitespace()

	switch l.ch {
	case '=':
		// ==
		if l.peekChar() == '=' {
			// readCharを呼ぶ前に現在の文字を保持
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		// !=
		if l.peekChar() == '=' {
			// readCharを呼ぶ前に現在の文字を保持
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '"':
		// 文字列
		tok.Type = token.STRING
		// 次の '"'が現れるまでトークンを進めてその間の文字列を読み取る
		tok.Literal = l.readString()
	case 0:
		// 終端として扱う
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// l.chが認識された文字でないときに識別子かどうかを点検する
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			// 対象の文字をどのようにして扱えばいいかわからない場合
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}