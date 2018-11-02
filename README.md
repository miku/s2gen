# solrstructgen

Generate code for representing SOLR documents in Go.

Solr allows to specify a schema, consisting of fields, types, cardinality and
dynamic types. To facilitate the creation of documents conforming to a SOLR
schema, it would be nice, to generate a struct and code for a schema
automatically from XML.

The SOLR schema is documented at [Documents, Fields, and Schema
Design](https://lucene.apache.org/solr/guide/6_6/documents-fields-and-schema-design.html).

The following features are implemented:

* static fields with types and cardinality
* dynamic fields
