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

	"github.com/miku/solrstructgen"
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

var (
	SetTmpl = `func (%s *%s) Set(k, v string) error {
		return nil
	}`
	MarshalTmpl = `func(%s *%s) MarshalJSON() ([]byte, error) {
		return nil, nil
	}`
)

// GoName converts a string into a more idiomatic name. Might miss edge cases.
func GoName(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "-", "", -1)
	parts := strings.Split(s, "_")
	var camel []string
	for _, p := range parts {
		camel = append(camel, strings.Title(p))
	}
	return strings.Join(camel, "")
}

func main() {
	flag.Parse()

	var r io.Reader = os.Stdin
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

	var schema solrstructgen.Schema
	if err := dec.Decode(&schema); err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	io.WriteString(&buf, "type ")

	if schema.Name == "" {
		log.Fatal("schema does not has a name")
	}

	io.WriteString(&buf, GoName(schema.Name))
	io.WriteString(&buf, " struct {\n")

	for _, f := range schema.Fields.Field {
		log.Println(f.Name, f.Type)
		if f.MultiValued == "true" {
			fmt.Fprintf(&buf, "%s []string\n", GoName(f.Name))
		} else {
			fmt.Fprintf(&buf, "%s string\n", GoName(f.Name))
		}
		// XXX: Struct with normal fields.
		// XXX: Methods to add dynamic fields with checks, e.g. doc.Set("field", "value")
		// XXX: Custom JSON marshaller.
	}
	log.Printf("The %v %v schema contains %d static fields.\n",
		schema.Name, schema.Version, len(schema.Fields.Field))

	io.WriteString(&buf, "}")

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
