// Package uml is a extension for the goldmark(http://github.com/yuin/goldmark).
//
// This extension adds svg picture output from uml language using
// go-plantuml(https://github.com/OhYee/go-plantuml).
package uml

import (
	"bytes"
	"crypto/sha1"

	gouml "github.com/OhYee/go-plantuml"
	ext "github.com/OhYee/goldmark-fenced_codeblock_extension"
	fp "github.com/OhYee/goutils/functional"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

// Default uml extension when there is no other fencedCodeBlock goldmark render extensions
var Default = NewUMLExtension(50, "plantuml")

// RenderMap return the goldmark-fenced_codeblock_extension.RenderMap
func RenderMap(length int, languages ...string) ext.RenderMap {
	return ext.RenderMap{
		Languages:      languages,
		RenderFunction: NewUML(length, languages...).Renderer,
	}
}

// NewUMLExtension return the goldmark.Extender
func NewUMLExtension(length int, languages ...string) goldmark.Extender {
	return ext.NewExt(RenderMap(length, languages...))
}

// UML render struct
type UML struct {
	Languages []string
	buf       map[string][]byte
	MaxLength int
}

// NewUML initial a UML struct
func NewUML(length int, languages ...string) *UML {
	return &UML{Languages: languages, buf: make(map[string][]byte), MaxLength: length}
}

// Renderer render function
func (u *UML) Renderer(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	language := string(n.Language(source))

	if fp.AnyString(func(l string) bool {
		return l == language
	}, u.Languages) {
		if !entering {
			raw := u.getLines(source, node)
			h := sha1.New()
			h.Write(raw)
			hash := string(h.Sum([]byte{}))
			if result, exist := u.buf[hash]; exist {
				w.Write(result)
			} else {
				svg, _ := gouml.UML(raw)
				if len(u.buf) >= u.MaxLength {
					u.buf = make(map[string][]byte)
				}
				u.buf[hash] = svg
				w.Write(svg)
			}
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
