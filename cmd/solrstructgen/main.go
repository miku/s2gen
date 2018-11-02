// solrstructgen creates a Go struct that can be serialized to JSON, which can
// be indexed into SOLR. It is useful for the following scenario: You have
// input data in some (weird) format and you have a target format. Generate
// structs for both input and output and all that is left to do is to write
// some wiring of the fields.
package main

import (
	"bytes"
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

// XXX: This is just a snippet for reference.
var t = `
type SomeSchema struct {
	Title          string
	ContainerTitle string

	dfields []struct {
		Key   string
		Value string
	}
}

func (s *SomeSchema) Set(key, value string) error {
	// Check if key is static or dynamic.
	// Check for dynamic key regex contraints.
	// Add value to dfields.
}

func (s *SomeSchema) MarshalJSON() ([]byte, error) {
	// Marshal dynamic fields into normal fields.
}
`

// Methods template for dynamic field support.
var Methods = `
func ({{ .Var }} {{ .Name }}) IsValidDynamicFieldName(k string) bool {
}

// Set sets the value for a dynamic field, only. The key is validated against
// the dynamic field wildcard (https://is.gd/qD1d1N).
func ({{ .Var }} {{ .Name }}) Set(k, v string) error {
	return nil
}

// MarshalJSON serializes static and dynamic fields.
func({{ .Var }} {{ .Name }}) MarshalJSON() ([]byte, error) {
	return nil, nil
}

`

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

	dec := xml.NewDecoder(r)
	dec.Strict = false

	var schema ssg.Schema
	if err := dec.Decode(&schema); err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	io.WriteString(&buf, "type ")

	if schema.Name == "" {
		log.Fatal("schema does not has a name")
	}

	// Fix name of type and variable name.
	typeName := ssg.GoName(schema.Name)
	if typeName == "" {
		log.Fatal("the go name reduced to the empty string")
	}
	varName := strings.ToLower(typeName[0:1])

	io.WriteString(&buf, ssg.GoName(schema.Name))
	io.WriteString(&buf, " struct {\n")

	for _, f := range schema.Fields.Field {
		log.Println(f.Name, f.Type)
		if f.MultiValued == "true" {
			fmt.Fprintf(&buf, "%s []string\n", ssg.GoName(f.Name))
		} else {
			fmt.Fprintf(&buf, "%s string\n", ssg.GoName(f.Name))
		}
		// XXX: Struct with normal fields.
		// XXX: Methods to add dynamic fields with checks, e.g. doc.Set("field", "value")
		// XXX: Custom JSON marshaller.
	}
	log.Printf("The %v %v schema contains %d static fields.\n",
		schema.Name, schema.Version, len(schema.Fields.Field))

	io.WriteString(&buf, `
	dynamicFields []struct {
		Key    string
		Values []string
	}
	`)

	io.WriteString(&buf, "}")

	var dnames []string
	for _, f := range schema.Fields.DynamicField {
		dnames = append(dnames, f.Name)
	}

	mtmpl := `
	// allowedDynamicFieldName return true, if the name of the field matches
	// one of the dynamic field patterns.
	func ({{ .Var }} {{ .Name }}) allowedDynamicFieldName(k string) (ok bool, err error) {
		return WildcardMatch(k, {{ .NameSlice }})
	}

	// WildcardMatch returns true, if the wildcard covers a given string s. If the
	// wildcard is invalid, an error is returned.
	func WildcardMatch(s, wildcards []string) (bool, error) {
		for _, w := range wildcards {
			p := strings.Replace(w, "*", ".*", -1)
			p = "^" + p + "$"
			r, err := regexp.Compile(p)
			if err != nil {
				return false, err
			}
			if r.MatchString(s) {
				return true, nil
			}
		}
		return false, nil
	}
	`

	t := template.New("methods")
	t, err := t.Parse(mtmpl)
	if err != nil {
		log.Fatal(err)
	}

	var data = struct {
		Var       string
		Name      string
		NameSlice string
	}{
		Var: varName, Name: typeName, NameSlice: ssg.RenderStringSlice(dnames),
	}

	if err := t.Execute(&buf, data); err != nil {
		log.Fatal(err)
	}

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
