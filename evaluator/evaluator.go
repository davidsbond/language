// Package evaluator contains methods for traversing the abstract syntax tree.
// They convert literals into objects in memory & manage the scope for each
// node.
package evaluator

import (
	"github.com/davidsbond/dave/ast"
	"github.com/davidsbond/dave/object"
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
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.NumberLiteral:
		return &object.Number{Value: node.Value}
	case *ast.BooleanLiteral:
		return &object.Boolean{Value: node.Value}
	case *ast.CharacterLiteral:
		return &object.Character{Value: node.Value}
	case *ast.Identifier:
		return scope.Get(node.Value)

	}

	return nil
}

func evaluateAST(ast *ast.AST, scope *object.Scope) object.Object {
	var result object.Object

	for _, node := range ast.Nodes {
		result = Evaluate(node, scope)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		}
	}

	return result
}
