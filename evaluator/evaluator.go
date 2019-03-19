package evaluator
import (
	"fmt"

	"github.com/kakts/monkey/ast"
	"github.com/kakts/monkey/object"
)

// Null Object.Booolean値をあらかじめ生成して参照するようにする
var (
	NULL = &object.Null{}
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
	case *ast.PrefixExpression:
		// 前置詞
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		// 中置
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)

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

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return NULL
	}
}

// !前置詞を含む場合の評価
func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

// -前置詞を含む場合の評価
func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return NULL
	}

	value := right.(*object.Integer).Value
	// 正負を反転した上で整数オブジェクトのインスタンスを返す
	return &object.Integer{Value: -value}
}

// 中置式の評価
func evalInfixExpression (
	operator string,
	left, right object.Object,
) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		// 左右どちらも整数の場合
		return evalIntegerInfixExpression(operator, left, right)
	default:
		return NULL
	}
}

// 左右が整数の場合の中置式の評価
func evalIntegerInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		// ゼロ除算考慮しない
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return NULL
	}
}