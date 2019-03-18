package evaluator
import (
	"fmt"

	"github.com/kakts/monkey/ast"
	"github.com/kakts/monkey/object"
)

// ASTノードを評価する
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		// 文
		fmt.Println("*ast.Program")
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		// 式 再帰的に評価
		fmt.Println("*ast.ExpressionStatement")
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return &object.Boolean{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}