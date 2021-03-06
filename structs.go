package main

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/types"
)

// Library represents a collection of classes, variables, and functions.
type Library struct {
	Name       string
	Classes    map[string]*Class
	Interfaces []*ast.GenDecl
	FuncTypes  []*ast.GenDecl
	Funcs      []*ast.FuncDecl
	Vars       []*ast.GenDecl
	Types      *types.Info
	Imports    []*ast.ImportSpec
	Files      []string
	CommentMap []ast.CommentMap
	Fset       *token.FileSet
}

func NewLibrary() *Library {
	return &Library{
		Name:       "",
		Classes:    map[string]*Class{},
		Interfaces: []*ast.GenDecl{},
		Funcs:      []*ast.FuncDecl{},
		Vars:       []*ast.GenDecl{},
		Types:      nil,
		Imports:    []*ast.ImportSpec{},
	}
}

// Class represents a dart class
type Class struct {
	Name    string
	Fields  []*ast.Field
	Methods []*ast.FuncDecl
}
