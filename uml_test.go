package uml

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"testing"

	gouml "github.com/OhYee/go-plantuml"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

func Test_default(t *testing.T) {
	var buf bytes.Buffer
	source := []byte("```go\npackage main\n\nimport ()\n\nfunc main(){}\n```\n\n```plantuml\n@startuml\nAlice -> Bob: test\n@enduml\n```\n\n")
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

	buf.Reset()

	if err := md.Convert(source, &buf); err != nil {
		t.Error(err)
	}
	if bytes.Compare(buf.Bytes(), []byte(want)) != 0 {
		t.Errorf("got %s, excepted %s\n", buf.Bytes(), want)
	}
}

func Test_demo(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	goPath, err := exec.LookPath("go")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	demoDir := path.Join(path.Dir(file), "demo")
	var cmd *exec.Cmd

	os.Remove(path.Join(demoDir, "demo1", "output.html"))
	cmd = exec.Command(goPath, "run", path.Join(demoDir, "demo1", "main.go"))
	if err := cmd.Run(); err != nil {
		t.Errorf("Error: %+v", err)
		t.FailNow()
	}
	if data, err := ioutil.ReadFile(path.Join(demoDir, "demo1", "output.html")); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		if c := strings.Count(string(data), "svg"); c != 3 {
			t.Errorf("Find %d svg", c)
			t.FailNow()
		}
	}

	os.Remove(path.Join(demoDir, "demo2", "output.html"))
	cmd = exec.Command(goPath, "run", path.Join(demoDir, "demo2", "main.go"))
	if err := cmd.Run(); err != nil {
		t.Errorf("Error: %+v", err)
		t.FailNow()
	}
	if data, err := ioutil.ReadFile(path.Join(demoDir, "demo2", "output.html")); err != nil {
		t.Error(err)
		t.FailNow()
	} else {
		if c := strings.Count(string(data), "svg"); c != 3 {
			t.Errorf("Find %d svg", c)
			t.FailNow()
		}
		if c := strings.Count(string(data), "color"); c != 5 {
			t.Errorf("Find %d color", c)
			t.FailNow()
		}
	}
}
