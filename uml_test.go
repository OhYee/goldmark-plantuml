package uml

import (
	"bytes"
	"testing"

	gouml "github.com/OhYee/go-plantuml"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

func Test_default(t *testing.T) {
	var buf bytes.Buffer
	source := []byte("```go\npackage main\n\nimport ()\n\nfunc main(){}\n```\n\n```uml\n@startuml\nAlice -> Bob: test\n@enduml\n```\n\n")
	want := `<pre><code class="language-go">package main

import ()

func main(){}
</code></pre>
` + func() string { a, _ := gouml.UML([]byte("@startuml\nAlice -> Bob: test\n@enduml")); return string(a) }()
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			Default,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(),
	)

	if err := md.Convert(source, &buf); err != nil {
		t.Error(err)
	}
	if bytes.Compare(buf.Bytes(), []byte(want)) != 0 {
		t.Errorf("got %s, excepted %s\n", buf.Bytes(), want)
	}
}
