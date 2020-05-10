// Package uml is a extension for the goldmark(http://github.com/yuin/goldmark).
//
// This extension adds svg picture output from uml language using
// go-plantuml(https://github.com/OhYee/go-plantuml).
package uml

import (
	"bytes"

	gouml "github.com/OhYee/go-plantuml"
	ext "github.com/OhYee/goldmark-fenced_codeblock_extension"
	fp "github.com/OhYee/goutils/functional"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

// Default uml extension when there is no other fencedCodeBlock goldmark render extensions
var Default = NewUMLExtension("plantuml")

// RenderMap return the goldmark-fenced_codeblock_extension.RenderMap
func RenderMap(languages ...string) ext.RenderMap {
	return ext.RenderMap{
		Languages:      languages,
		RenderFunction: NewUML(languages).Renderer,
	}
}

// NewUMLExtension return the goldmark.Extender
func NewUMLExtension(languages ...string) goldmark.Extender {
	return ext.NewExt(RenderMap(languages...))
}

// UML render struct
type UML struct {
	Languages []string
}

// NewUML initial a UML struct
func NewUML(languages []string) *UML {
	return &UML{languages}
}

// Renderer render function
func (u *UML) Renderer(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	language := string(n.Language(source))

	if fp.AnyString(func(l string) bool {
		return l == language
	}, u.Languages) {
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
