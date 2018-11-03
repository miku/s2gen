# solrstructgen

Generate code to represent SOLR documents in Go.

Solr allows to specify a
[schema](https://lucene.apache.org/solr/guide/6_6/documents-fields-and-schema-design.html),
consisting of fields, types, cardinality and dynamic types. To facilitate the
creation of documents conforming to a SOLR schema, it would be nice, to
generate a struct and code for a schema automatically from the XML schema definition.

The SOLR schema is documented at [Documents, Fields, and Schema
Design](https://lucene.apache.org/solr/guide/6_6/documents-fields-and-schema-design.html).

The following features are implemented:

* generating a struct from a SOLR schema.xml
* generating a indexable JSON document from a struct
* parsing SOLR documents into a struct
* static fields with types and cardinality
* dynamic fields

Not supported or low priority:

* type inference (to many variations, so we use strings for now)

## Usage

First, generate the code from the SOLR schema.

```shell
$ solrstructgen < schema.xml > schema.go
```

As an example, let say this generated a struct named `VuFindBibliographicIndex`
and a few helper functions. From here, you have various ways to populate the
struct values.

```go
var doc VuFindBibliographicIndex

// Manipulate static fields, multi-valued fields are slices.
doc.Fulltext = "This is the full text"
doc.Author = append(doc.Author, "Samuel Johnson")

// Set a single dynamic field. Error, if dynamic field is not valid.
err := doc.Set("format_de15", "Book")

// Set values for a bunch of dynamic fields at once.
err := doc.SetMap("format_*", map[string]interface{
    "de_15": "Book",
    "de_14": "Book",
})

// Create an indexable document.
b, _ := json.Marshal(doc)
fmt.Println(string(b))
```

## Use case

The use case is to allow for small converter programs for various formats:

* generate input XML or JSON struct
* generate target XML, JSON or SOLR struct

Then write a single function (should take less than one hour):

```go
func main() {
    var s SourceDoc // e.g. XML
    var t TargetDoc // e.g. SOLR

    t.Field = s.Field
    // Lookup, cleanup, filters.

    b, _ := json.Marshal(t)
    fmt.Println(string(b))
}
```

Lookup tables should be aided by function to fetch mappings and tables from
files, URLs, repos and more.

Another use case is the parsing and manipulation of SOLR documents in general:
fetch, parse, modify, serialize, index. For example, you want to rewrite a
subset of documents and modify a single field, then reindex.

For high-performance modifications, use smaller, focussed structs (with only
one or two relevant fields).

```go
type Doc struct {
    ID          string
    ISSN        []string
    Institution []string
    ...
}
```

* Parallel fetch, modification, indexing.


## TODO

* [ ] schema.xml -> dynamic fields
* [ ] marshal dynamic fields, MarshalJSON
* [ ] unmarshal dynamic fields, UnmarshalJSON
* [ ] tests

