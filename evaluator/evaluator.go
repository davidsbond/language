// Package evaluator contains methods for traversing the abstract syntax tree.
// They convert literals into objects in memory & manage the scope for each
// node.
package evaluator

import (
	"math"

	"github.com/davidsbond/language/ast"
	"github.com/davidsbond/language/builtin"
	"github.com/davidsbond/language/object"
)

var (
	builtins = map[string]object.Builtin{
		// Reflection builtins
		"type": builtin.Type,
		"len":  builtin.Len,

		// Environment builtins
		"set_env": builtin.SetEnv,
		"get_env": builtin.GetEnv,

		// Mathematical builtins
		"cos":   builtin.Cos,
		"tan":   builtin.Tan,
		"sin":   builtin.Sin,
		"log":   builtin.Log,
		"ceil":  builtin.Ceil,
		"floor": builtin.Floor,
	}

	values = map[string]object.Object{
		"π": &object.Number{Value: math.Pi},
		"∞": &object.Number{Value: math.Inf(0)},
	}

	// NULL is used as the global null object
	NULL = &object.Null{}

	// TRUE is used as the global true value
	TRUE = &object.Boolean{Value: true}

	// FALSE is used as the global false value
	FALSE = &object.Boolean{Value: false}
)

// Evaluate attempts to  the given node using the provided scope and return
// an object.
func Evaluate(node ast.Node, scope *object.Scope) object.Object {
	switch node := node.(type) {
	case *ast.AST:
		return evalAST(node, scope)
	case *ast.ExpressionStatement:
		return Evaluate(node.Expression, scope)
	case *ast.InfixExpression:
		return infixExpression(node, scope)
	case *ast.VarStatement:
		return varStatement(node, scope)
	case *ast.ConstStatement:
		return constStatement(node, scope)
	case *ast.AtomicStatement:
		return atomicStatement(node, scope)
	case *ast.BlockStatement:
		return blockStatement(node, scope)
	case *ast.CallExpression:
		return callExpression(node, scope)
	case *ast.PrefixExpression:
		return prefixExpression(node, scope)
	case *ast.IndexExpression:
		return indexExpression(node, scope)
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.NumberLiteral:
		return &object.Number{Value: node.Value}
	case *ast.CharacterLiteral:
		return &object.Character{Value: node.Value}
	case *ast.FunctionLiteral:
		return functionLiteral(node, scope)
	case *ast.ReturnStatement:
		return &object.ReturnValue{Value: Evaluate(node.ReturnValue, scope)}
	case *ast.Identifier:
		return identifier(node, scope)
	case *ast.AsyncStatement:
		return asyncStatement(node, scope)
	case *ast.AwaitStatement:
		return awaitStatement(node, scope)
	case *ast.ArrayLiteral:
		return arrayLiteral(node, scope)
	case *ast.HashLiteral:
		return hashLiteral(node, scope)
	case *ast.PostfixExpression:
		return postfixExpression(node, scope)
	case *ast.IfExpression:
		return ifExpression(node, scope)
	case *ast.Comment:
		return nil
	case *ast.BooleanLiteral:
		if node.Value {
			return TRUE
		}

		return FALSE
	case nil:
		return object.Error("cannot  nil node")
	default:
		return object.Error("unsupported node type %s", node.String())
	}
}

func evalAST(ast *ast.AST, scope *object.Scope) object.Object {
	var result object.Object

	for _, node := range ast.Nodes {
		result = Evaluate(node, scope)

		if result != nil && result.Type() == object.TypeError {
			break
		}

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		}
	}

	return result
}

func isError(obj object.Object) bool {
	return obj.Type() == object.TypeError
}
