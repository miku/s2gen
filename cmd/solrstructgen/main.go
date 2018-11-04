// solrstructgen creates a Go struct that can be serialized to JSON, which can
// be indexed into SOLR. It is useful for the following scenario: You have
// input data in some (weird) format and you have a target format. Generate
// structs for both input and output and all that is left to do is to write
// some wiring of the fields.
package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/xml"
	"flag"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	ssg "github.com/miku/solrstructgen"
)

var (
	skipFormatting = flag.Bool("F", false, "skip formatting")
)

func main() {
	var r io.Reader = os.Stdin

	flag.Parse()
	if flag.NArg() > 0 {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		r = f
	}

	// Calculate a hash of the content.
	h := sha1.New()
	tee := io.TeeReader(r, h)

	// Decode schema.
	dec := xml.NewDecoder(tee)
	dec.Strict = false

	var schema ssg.Schema
	if err := dec.Decode(&schema); err != nil {
		log.Fatal(err)
	}

	digest := fmt.Sprintf("%x", h.Sum(nil))

	var buf bytes.Buffer

	io.WriteString(&buf, `
	package main

	import (
		"encoding/json"
		"fmt"
		"log"
		"os"
		"regexp"
		"strings"
	)
	`)

	// Fix name of type and variable name.
	typeName := ssg.GoName(schema.Name)
	if typeName == "" {
		log.Fatal("the go name reduced to the empty string")
	}
	// Name to use for struct methods.
	varName := strings.ToLower(typeName[0:1])

	fmt.Fprintf(&buf, "// %s generated from %s.\n", typeName, digest)
	io.WriteString(&buf, "type ")

	if schema.Name == "" {
		log.Fatal("schema has no name")
	}

	io.WriteString(&buf, ssg.GoName(schema.Name))
	io.WriteString(&buf, " struct {\n")

	for _, f := range schema.Fields.Field {
		switch {
		case f.Name == "_version_":
			fmt.Fprintf(&buf, "%s json.Number `json:\"%s\"`\n", ssg.GoName(f.Name), f.Name)
		case f.MultiValued == "true":
			fmt.Fprintf(&buf, "%s []string `json:\"%s\"`\n", ssg.GoName(f.Name), f.Name)
		default:
			fmt.Fprintf(&buf, "%s string `json:\"%s\"`\n", ssg.GoName(f.Name), f.Name)
		}
		// XXX: Methods to add dynamic fields with checks, e.g. doc.Set("field", "value")
		// XXX: Custom JSON marshaller.
	}

	io.WriteString(&buf, `
	dynamicFields []DynamicField
	`)

	io.WriteString(&buf, "}")

	// Support and struct methods.
	var dnames []string
	for _, f := range schema.Fields.DynamicField {
		dnames = append(dnames, f.Name)
	}

	mtmpl := `
	type DynamicField struct {
		Name        string
		MultiValued bool
		Values      map[string][]string
	}

	// allowedDynamicFieldName returns true, if the name of the field matches
	// one of the dynamic field patterns.
	func ({{ .Var }} {{ .Name }}) allowedDynamicFieldName(k string) error {
		return WildcardMatch(k, {{ .NameSlice }})
	}

	// WildcardMatch returns nil, if the wildcard covers a given string s. If the
	// wildcard is invalid, an error is returned.
	func WildcardMatch(s string, wildcards []string) error {
		// XXX: It is possible, that a static field will match a dynamic field.
		for _, w := range wildcards {
			p := strings.Replace(w, "*", ".*", -1)
			p = "^" + p + "$"
			r, err := regexp.Compile(p)
			if err != nil {
				return err
			}
			if r.MatchString(s) {
				return nil
			}
		}
		return fmt.Errorf("wildcards do not cover key: %s", s)
	}

	func main() {
		var doc {{ .Name }}

		dec := json.NewDecoder(os.Stdin)
		if err := dec.Decode(&doc); err != nil {
			log.Fatal(err)
		}

		fmt.Println(doc)
	}
	`

	// Render template.
	// XXX: It is possible, that a static field will match a dynamic field.
	t, err := template.New("methods").Parse(mtmpl)
	if err != nil {
		log.Fatal(err)
	}

	var data = struct {
		Var       string
		Name      string
		NameSlice string
	}{
		Var:       varName,
		Name:      typeName,
		NameSlice: ssg.RenderStringSlice(dnames),
	}

	if err := t.Execute(&buf, data); err != nil {
		log.Fatal(err)
	}

	// Format output.
	if *skipFormatting {
		fmt.Println(buf.String())
	} else {
		b, err := format.Source(buf.Bytes())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	}
}
