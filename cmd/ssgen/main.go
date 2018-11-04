// ssgen takes as input a SOLR schema.xml and outputs structs and methods to
// access documents conforming to that schema.
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
	"os/user"
	"strings"
	"text/template"
	"time"

	ssg "github.com/miku/solrstructgen"
)

const Version = "0.1.0"

var (
	skipFormatting   = flag.Bool("F", false, "skip formatting")
	showVersion      = flag.Bool("version", false, "show version")
	useHashSuffix    = flag.Bool("hs", false, "add part of hash as suffix to struct name")
	useVersionSuffix = flag.String("vs", "", "add some versioning suffix, like v0 to struct name")
)

// field contains information about static fields.
type field struct {
	GoName string
	GoType string
	GoTag  string
}

// dynamicField contains basic information on a dynamic field
type dynamicField struct {
	Name          string
	IsMultiValued string
}

// payload provides data for template.
type payload struct {
	Name          string
	VarName       string
	Digest        string
	Date          string
	User          string
	Host          string
	Version       string
	Fields        []field
	DynamicFields []dynamicField
}

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	var r io.Reader = os.Stdin

	if flag.NArg() > 0 {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		r = f
	}

	h := sha1.New()
	tee := io.TeeReader(r, h)

	dec := xml.NewDecoder(tee)
	dec.Strict = false

	var schema ssg.Schema
	if err := dec.Decode(&schema); err != nil {
		log.Fatal(err)
	}

	digest := fmt.Sprintf("%x", h.Sum(nil))

	// Fix name of type and variable name.
	name := ssg.GoName(schema.Name)
	if name == "" {
		log.Fatal("the go name reduced to the empty string")
	}
	varName := strings.ToLower(name[0:1])

	// Optional suffixes.
	if *useHashSuffix {
		name = name + digest[0:8]
	}
	if *useVersionSuffix != "" {
		name = name + *useVersionSuffix
	}

	// Some metadata.
	usrName := "an unknown user"
	usr, _ := user.Current()
	if usr != nil {
		usrName = usr.Username
	}
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "an unknown host"
	}

	// Payload for template.
	data := payload{
		Name:    name,
		VarName: varName,
		Digest:  digest,
		Date:    time.Now().Format(time.RFC3339),
		User:    usrName,
		Host:    hostname,
		Version: Version,
	}

	for _, f := range schema.Fields.Field {
		ff := field{
			GoName: ssg.GoName(f.Name),
			GoTag:  fmt.Sprintf("`json:\"%s\"`", f.Name),
		}
		switch {
		case f.Name == "_version_":
			ff.GoType = "json.Number"
		case f.MultiValued == "true":
			ff.GoType = "[]string"
		default:
			ff.GoType = "string"
		}
		data.Fields = append(data.Fields, ff)
	}

	for _, f := range schema.Fields.DynamicField {
		if f.MultiValued == "" {
			f.MultiValued = "false"
		}
		ff := dynamicField{Name: f.Name, IsMultiValued: f.MultiValued}
		data.DynamicFields = append(data.DynamicFields, ff)
	}

	// Render template.
	t, err := template.New("document.tmpl").ParseFiles("tmpl/document.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
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
