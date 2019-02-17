package parser

import (
	"github.com/kakts/monkey/ast"
	"github.com/kakts/monkey/lexer"
	"github.com/kakts/monkey/token"
)

type Parser struct {
	l *lexer.lexer

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// 2つのトークンを読み込む
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}