package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/antlr4-go/antlr/v4"

	"github.com/sjansen/shiny/spikes/partiql/parser"
)

type StdinStream struct {
	*antlr.InputStream
}

func NewStdinStream() (*StdinStream, error) {
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, os.Stdin)
	if err != nil {
		return nil, err
	}

	fs := new(StdinStream)
	fs.InputStream = antlr.NewInputStream(
		buf.String(),
	)

	return fs, nil
}

func (f *StdinStream) GetSourceName() string {
	return "<STDIN>"
}

type TreeShapeListener struct {
	*parser.BasePartiQLListener
	depth int
}

func NewTreeShapeListener() *TreeShapeListener {
	return &TreeShapeListener{}
}

func (l *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("%*s %T    %s\n", l.depth, "", ctx, ctx.GetText())
	l.depth++
}

func (l *TreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	l.depth--
}

func main() {
	var err error
	var input antlr.CharStream
	if len(os.Args) < 2 || os.Args[1] == "-" {
		input, err = NewStdinStream()
	} else {
		input, err = antlr.NewFileStream(os.Args[1])
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	lexer := parser.NewPartiQLTokens(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPartiQLParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Root()
	fmt.Println("=========================================")
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
