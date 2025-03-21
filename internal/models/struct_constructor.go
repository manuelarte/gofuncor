package models

import (
	"go/ast"

	"github.com/manuelarte/gofuncor/internal/utils"
)

type StructConstructor struct {
	constructor  *ast.FuncDecl
	structReturn *ast.Ident
}

func NewStructConstructor(funcDec *ast.FuncDecl) (StructConstructor, bool) {
	if utils.FuncCanBeConstructor(funcDec) {
		expr := funcDec.Type.Results.List[0].Type
		if returnType, ok := utils.GetIdent(expr); ok {
			return StructConstructor{
				constructor:  funcDec,
				structReturn: returnType,
			}, true
		}
	}
	return StructConstructor{}, false
}

// GetStructReturn Return the struct linked to this "constructor".
func (sc StructConstructor) GetStructReturn() *ast.Ident {
	return sc.structReturn
}

func (sc StructConstructor) GetConstructor() *ast.FuncDecl {
	return sc.constructor
}
