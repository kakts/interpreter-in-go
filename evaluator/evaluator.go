package evaluator
import (
	"fmt"

	"github.com/kakts/monkey/ast"
	"github.com/kakts/monkey/object"
)

// Object.Booolean値をあらかじめ生成して参照するようにする
var (
	TRUE = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
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
		// プリミティブ値からBooleanオブジェクトのインスタンスを取得する
		return nativeBoolToBooleanObject(node.Value)
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

// プリミティブ値からBooleanオブジェクトのインスタンスを取得する
func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}