// ssgen takes as input a SOLR schema.xml and outputs structs and methods to
// access documents conforming to that schema.
package main

import (
	"crypto/sha1"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	skipFormatting = flag.Bool("F", false, "skip formatting")
)

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

	hash := sha1.New()
	tee := io.TeeReader(r, h)

	dec := xml.NewDecoder(tee)
	dec.Strict = false

	var schema ssg.Schema
	if err := dec.Decode(&schema); err != nil {
		log.Fatal(err)
	}

	digest := fmt.Sprintf("%x", hash.Sum(nil))
	fmt.Println(digest)
}
