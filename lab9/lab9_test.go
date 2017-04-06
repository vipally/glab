package lab9_test

import (
	"fmt"
	"go/parser"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"testing"
)

func TestGo(t *testing.T) {
	file_ := "E:\\dev\\gocode\\trunk\\src\\git.oschina.net\\vipally\\gx\\regable\\gp\\factoryfix.gp.go"
	src, _ := ioutil.ReadFile(file_)
	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Sprintf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
	// Parse the file containing this very example
	// but stop after processing the imports.
	f, err := parser.ParseFile(fset, file_, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the imports from the file's AST.
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}
	for _, s := range f.Comments {
		fmt.Println(s.Pos(), s.Text())
	}
	for _, s := range f.Decls {
		fmt.Println(s.Pos(), s.End())
	}
	for _, s := range f.Unresolved {
		fmt.Println(s.String())
	}
	fmt.Printf("%#v\n %#v\n %#v\n %#v\n %#v\n", f.Doc, f.Scope, f.Name, f.Package, f.Unresolved)
}
