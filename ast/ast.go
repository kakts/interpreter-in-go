package ast
import (
	"github.com/kakts/monkey/token"
)
// ASTのすべてのノードが実装しなければならない つまり TokenLiteral()を実装する必要がある
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// 構文解析器が生成するすべてのASTのルートノードになるもの
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) >  0 {
		// 先頭statementの処理
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}


type LetStatement struct {
	Token token.Token // token.LET トークン
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}