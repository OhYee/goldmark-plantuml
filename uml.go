// Package uml is a extension for the goldmark(http://github.com/yuin/goldmark).
//
// This extension adds svg picture output from uml language using
// gouml(https://github.com/OhYee/gouml).
package uml

import (
	"bytes"

	gouml "github.com/OhYee/go-plantuml"
	ext "github.com/OhYee/goldmark-fenced_codeblock_extension"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

// Default  uml extension when there is no other fencedCodeBlock goldmark render extensions
var Default = NewUMLExtension("uml")

func NewUMLExtension(languageName string) goldmark.Extender {
	return ext.NewExt([]ext.RenderMap{
		ext.RenderMap{
			Language:       []string{languageName},
			RenderFunction: NewUML(languageName).Renderer,
		},
	}...)
}

type UML struct {
	LanguageName string
}

func NewUML(languageName string) *UML {
	return &UML{languageName}
}

func (u *UML) Renderer(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	language := n.Language(source)
	if string(language) == u.LanguageName {
		if !entering {
			svg, _ := gouml.UML(u.getLines(source, node))
			w.Write(svg)
		}
	}
	return ast.WalkContinue, nil
}

func (u *UML) getLines(source []byte, n ast.Node) []byte {
	buf := bytes.NewBuffer([]byte{})
	l := n.Lines().Len()
	for i := 0; i < l; i++ {
		line := n.Lines().At(i)
		buf.Write(line.Value(source))
	}
	return buf.Bytes()
}
