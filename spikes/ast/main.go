package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strings"
	"unicode/utf8"
)

const directivePrefix = "shiny:"

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}
var reShinyDirective *regexp.Regexp

func directives(g *ast.CommentGroup) map[string]string {
	if reShinyDirective == nil {
		reShinyDirective = regexp.MustCompile(`^shiny:[a-z0-9]+`)
	}

	comments := make(map[string]string)
	if g == nil {
		return comments
	}
	for _, c := range g.List {
		text := c.Text
		if len(text) < 2 || text[0] != '/' {
			continue
		}
		switch text[1] {
		case '/':
			text = text[2:]
		case '*':
			text = strings.TrimSpace(text[2 : len(text)-2])

		}
		if text, ok := strings.CutPrefix(text, directivePrefix); ok {
			idx := 0
			for ; idx < len(text); idx++ {
				c := text[idx]
				if asciiSpace[c] != 0 || c >= utf8.RuneSelf {
					break
				}
			}
			key := text[:idx]
			value := strings.TrimSpace(text[idx:])
			comments[key] = value
		}
	}

	return comments
}

func listStructs(filename string) (map[string]map[string]string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	structs := make(map[string]map[string]string)
	ast.Inspect(file, func(n ast.Node) bool {
		if decl, ok := n.(*ast.GenDecl); ok && decl.Tok == token.TYPE {
			for _, spec := range decl.Specs {
				if ts, ok := spec.(*ast.TypeSpec); ok {
					if _, ok := ts.Type.(*ast.StructType); ok {
						name := ts.Name.Name
						structs[name] = directives(decl.Doc)
					}
				}
			}
		}
		return true
	})

	return structs, nil
}

func main() {
	// Test the function with a sample file
	structs, err := listStructs("testdata/sample.go")
	if err != nil {
		fmt.Println(err)
	} else {
		for name, comments := range structs {
			fmt.Println("---")
			fmt.Println(name)
			for directive, value := range comments {
				fmt.Printf("> %q %#v\n", directive, value)
			}
		}
	}
}
