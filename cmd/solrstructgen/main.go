// solrstructgen creates a Go struct that can be serialized to JSON, which can
// be indexed into SOLR. It is useful for the following scenario: You have
// input data in some (weird) format and you have a target format. Generate
// structs for both input and output and all that is left to do is to write
// some wiring of the fields.
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/miku/solrstructgen"
)

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

func main() {
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
	fmt.Println(schema.Version)
	for _, f := range schema.Fields.Field {
		fmt.Println(f.Name, f.Type)
		// XXX: Struct with normal fields.
		// XXX: Methods to add dynamic fields with checks, e.g. doc.Set("field", "value")
		// XXX: Custom JSON marshaller.
	}
	log.Printf("The %v %v schema contains %d static fields.\n",
		schema.Name, schema.Version, len(schema.Fields.Field))
}
