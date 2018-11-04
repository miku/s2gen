# s2gen (solr struct generator)

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
* marhsal an indexable JSON document from a struct
* unmarshal SOLR documents into a struct including dynamic fields
* static and dynamic fields with cardinality checks

Not supported or low priority:

* type inference (to many variations, so we use strings for now)

## Usage

First, generate the code from the SOLR schema.

```shell
$ s2gen < schema.xml > schema.go
```

As an example, let say this generated a struct named `VuFindBibliographicIndex`
and a few helper functions. From here, you have various ways to populate the
struct values.

```go
doc := NewVuFindBibliographicIndex()

// Manipulate static fields, multi-valued fields are slices.
doc.Fulltext = "This is the full text"
doc.Author = append(doc.Author, "Samuel Johnson")

// Set a single dynamic field. Error, if dynamic field is not valid.
err := doc.Set("format_de15", "Book")

// Create an indexable document.
b, _ := json.Marshal(doc)
fmt.Println(string(b))
```

## Use case

The use case is to allow for small converter programs for various formats:

* generate struct from input XML or JSON
* generate struct for target SOLR schema

Then write a single function (should take a few minutes):

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

* [x] schema.xml -> dynamic fields
* [x] marshal dynamic fields, MarshalJSON
* [x] unmarshal dynamic fields, UnmarshalJSON
* [ ] tests

