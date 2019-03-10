// Package evaluator contains methods for traversing the abstract syntax tree.
// They convert literals into objects in memory & manage the scope for each
// node.
package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/builtin"
	"github.com/davidsbond/dave/object"
)

var (
	builtins = map[string]object.Builtin{
		"type":    builtin.Type,
		"len":     builtin.Len,
		"set_env": builtin.SetEnv,
	}
)

// Evaluate attempts to evaluate the given node using the provided scope and return
// an object.
func Evaluate(node ast.Node, scope *object.Scope) object.Object {
	switch node := node.(type) {
	case *ast.AST:
		return evaluateAST(node, scope)
	case *ast.ExpressionStatement:
		return Evaluate(node.Expression, scope)
	case *ast.InfixExpression:
		return evaluateInfixExpression(node, scope)
	case *ast.VarStatement:
		return evaluateVarStatement(node, scope)
	case *ast.ConstStatement:
		return evaluateConstStatement(node, scope)
	case *ast.AtomicStatement:
		return evaluateAtomicStatement(node, scope)
	case *ast.BlockStatement:
		return evaluateBlockStatement(node, scope)
	case *ast.CallExpression:
		return evaluateCallExpression(node, scope)
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.NumberLiteral:
		return &object.Number{Value: node.Value}
	case *ast.BooleanLiteral:
		return &object.Boolean{Value: node.Value}
	case *ast.CharacterLiteral:
		return &object.Character{Value: node.Value}
	case *ast.FunctionLiteral:
		return evaluateFunctionLiteral(node, scope)
	case *ast.ReturnStatement:
		return &object.ReturnValue{Value: Evaluate(node.ReturnValue, scope)}
	case *ast.Identifier:
		return evaluateIdentifier(node, scope)
	case *ast.AsyncStatement:
		return evaluateAsyncStatement(node, scope)
	case *ast.AwaitStatement:
		return evaluateAwaitStatement(node, scope)
	case *ast.ArrayLiteral:
		return evaluateArrayLiteral(node, scope)
	case nil:
		return object.Error("cannot evaluate nil node")
	default:
		return object.Error("unsupported node type %s", node.String())
	}
}

func evaluateAST(ast *ast.AST, scope *object.Scope) object.Object {
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
