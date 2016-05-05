// Copyright 2016 orivil Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package comment provide one method for read the comments which in struct, method and function.
package comment

import (
	"go/parser"
	"go/token"
	"os"
	"go/doc"
)

type StructComment map[string]string		// {package.struct: comment}

type MethodComment map[string]map[string]string	// {package.struct: {methodName: comment}}

type FuncComment map[string]string		// {package.func: comment}

func GetDirComment(filter func(f os.FileInfo) bool, dirs... string) (sc StructComment, mc MethodComment, fc FuncComment, e error) {

	sc = make(StructComment, 2)
	mc = make(MethodComment, 5)
	fc = make(FuncComment, 2)
	for _, dir := range dirs {
		pkgs, err := parser.ParseDir(token.NewFileSet(), dir, filter, parser.ParseComments)
		if err != nil {
			e = err
			return
		}

		for _, p := range pkgs {
			doc := doc.New(p, dir, doc.AllDecls)

			for _, f := range doc.Funcs {
				fc[p.Name + "." + f.Name] = f.Doc
			}

			for _, t := range doc.Types {
				pkgStruct := p.Name + "." + t.Name
				sc[pkgStruct] = t.Doc
				if len := len(t.Methods); len > 0 {
					mc[pkgStruct] = make(map[string]string, len)
				}
				for _, m := range t.Methods {
					mc[pkgStruct][m.Name] = m.Doc
				}
			}
		}
	}
	return
}
