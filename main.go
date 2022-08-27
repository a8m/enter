package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	path := "./ent/schema"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	graph, err := entc.LoadGraph(path, &gen.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	var b bytes.Buffer
	if err := tmpl.Execute(&b, graph); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("er.html", b.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
}

var tmpl = template.Must(template.New("er").
	Funcs(template.FuncMap{
		"fmtType": func(s string) string {
			return strings.NewReplacer(
				".", "DOT",
				"*", "STAR",
				"[", "LBRACK",
				"]", "RBRACK",
			).Replace(s)
		},
	}).
	Parse(`<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
</head>
<body>
	{{- with $.Nodes }}
		<div class="mermaid" id="er-diagram">
erDiagram
{{- range $n := . }}
    {{ $n.Name }} {
	{{- if $n.HasOneFieldID }}
        {{ fmtType $n.ID.Type.String }} {{ $n.ID.Name }}
	{{- end }}
	{{- range $f := $n.Fields }}
        {{ fmtType $f.Type.String }} {{ $f.Name }}
	{{- end }}
    }
{{- end }}
{{- range $n := . }}
    {{- range $e := $n.Edges }}
	{{- if not $e.IsInverse }}
		{{- $rt := "|o--o|" }}{{ if $e.O2M }}{{ $rt = "|o--o{" }}{{ else if $e.M2O }}{{ $rt = "}o--o|" }}{{ else if $e.M2M }}{{ $rt = "}o--o{" }}{{ end }}
    	{{ $n.Name }} {{ $rt }} {{ $e.Type.Name }} : "{{ $e.Name }}{{ with $e.Ref }}/{{ .Name }}{{ end }}"
	{{- end }}
	{{- end }}
{{- end }}
		</div>
	{{- end }}
	<script src="https://cdn.jsdelivr.net/npm/mermaid/dist/mermaid.min.js"></script>
	<script>
		mermaid.mermaidAPI.initialize({
			startOnLoad: true,
		});
		var observer = new MutationObserver((event) => {
			document.querySelectorAll('text[id^=entity]').forEach(text => {
				text.textContent = text.textContent.replace('DOT', '.');
				text.textContent = text.textContent.replace('STAR', '*');
				text.textContent = text.textContent.replace('LBRACK', '[');
				text.textContent = text.textContent.replace('RBRACK', ']');
			});
			observer.disconnect();
		});
		observer.observe(document.getElementById('er-diagram'), { attributes: true, childList: true });
	</script>
</body>
</html>
`))
