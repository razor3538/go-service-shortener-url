package escape_os_exit

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"
)

var EscapeOsExit = &analysis.Analyzer{
	Name: "osexitcheck",
	Doc:  "check for os exit in main.go",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			if file.Name.String() == "main" {
				var funcs []*ast.FuncDecl

				for _, d := range file.Decls {
					if fn, isFn := d.(*ast.FuncDecl); isFn {
						funcs = append(funcs, fn)
					}
				}

				for _, function := range funcs {
					hasErr, pos := extractFuncCallInFunc(function.Body.List)

					if hasErr {
						pass.Reportf(pos, "os.Exit in main function")
						return false
					}
				}
			}
			return true
		})
	}

	return nil, nil
}

func extractFuncCallInFunc(stmts []ast.Stmt) (bool, token.Pos) {
	for _, stmt := range stmts {
		if exprStmt, ok := stmt.(*ast.ExprStmt); ok {
			if call, ok := exprStmt.X.(*ast.CallExpr); ok {
				if fun, ok := call.Fun.(*ast.SelectorExpr); ok {
					funcName := fun.Sel.Name
					if funcName == "Exit" {
						return true, fun.Sel.Pos()
					}
				}
			}
		}
	}
	return false, stmts[0].Pos()
}
