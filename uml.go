// package uml is a extension for the goldmark(http://github.com/yuin/goldmark).
//
// This extension adds svg picture output from uml language using
// gouml(https://github.com/OhYee/gouml).
package uml

import (
	"bytes"
	"fmt"

	gouml "github.com/OhYee/go-plantuml"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

// Config struct holds options for the extension.
type Config struct {
	languageName      string
	defaultRenderFunc renderer.NodeRendererFunc
}

type uml struct {
	config *Config
}

// UML default uml extension when there is no other fencedCodeBlock goldmark render extensions
var UML = NewUML("uml", html.NewRenderer())

// Newuml returns a new extension with given arguments.
func NewUML(languageName string, render renderer.NodeRenderer) goldmark.Extender {
	defaultRenderFunc := getRenderFunction(ast.KindFencedCodeBlock, render)
	if defaultRenderFunc == nil {
		panic(fmt.Sprintf("%T don't render ast.KindFencedCodeBlock(FencedCodeBlock)", render))
	}
	config := &Config{languageName, defaultRenderFunc}
	return &uml{config}
}

// Extend implements goldmark.Extender.
func (d *uml) Extend(m goldmark.Markdown) {
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewHTMLRenderer(d.config), 0),
	))
}

// HTMLRenderer struct is a renderer.NodeRenderer implementation for the extension.
type HTMLRenderer struct {
	config *Config
}

// NewHTMLRenderer builds a new HTMLRenderer with given options and returns it.
func NewHTMLRenderer(config *Config) renderer.NodeRenderer {
	r := &HTMLRenderer{config}
	return r
}

// RegisterFuncs implements NodeRenderer.RegisterFuncs.
func (r *HTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindFencedCodeBlock, r.renderFencedCodeBlock)
}

func (r *HTMLRenderer) getLines(source []byte, n ast.Node) []byte {
	buf := bytes.NewBuffer([]byte{})
	l := n.Lines().Len()
	for i := 0; i < l; i++ {
		line := n.Lines().At(i)
		buf.Write(line.Value(source))
	}
	return buf.Bytes()
}

func (r *HTMLRenderer) renderFencedCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	language := n.Language(source)

	if string(language) == r.config.languageName {
		if !entering {
			svg, _ := gouml.UML(r.getLines(source, node))
			w.Write(svg)
		}
	} else {
		return r.config.defaultRenderFunc(w, source, node, entering)
	}

	return ast.WalkContinue, nil
}

type hack struct {
	target   ast.NodeKind
	receiver *renderer.NodeRendererFunc
}

func (h hack) Register(node ast.NodeKind, f renderer.NodeRendererFunc) {
	if node.String() == h.target.String() {
		*h.receiver = f
	}
}

func getRenderFunction(target ast.NodeKind, r renderer.NodeRenderer) renderer.NodeRendererFunc {
	var receiver renderer.NodeRendererFunc
	h := hack{target, &receiver}
	r.RegisterFuncs(h)
	return receiver
}
