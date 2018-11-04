package s2gen

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// UppercaseWords contains lowercase strings, which should be uppercased for an
// idiomatic Go name.
var UppercaseWords = []string{
	"id",
	"iln",
	"isbn",
	"ismn",
	"issn",
	"json",
	"lccn",
	"marc",
	"marcxml",
	"oclc",
	"ppn",
	"rsn",
	"rvk",
	"uri",
	"url",
	"urn",
	"xml",
	"zdb",
}

func stringSliceContains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

// Schema was generated 2018-11-02 15:49:57 by tir on sol.
type Schema struct {
	XMLName xml.Name `xml:"schema"`
	Text    string   `xml:",chardata"`
	Name    string   `xml:"name,attr"`
	Version string   `xml:"version,attr"`
	Types   struct {
		Text      string `xml:",chardata"`
		FieldType []struct {
			Text                 string `xml:",chardata"`
			Name                 string `xml:"name,attr"`
			Class                string `xml:"class,attr"`
			PrecisionStep        string `xml:"precisionStep,attr"`
			PositionIncrementGap string `xml:"positionIncrementGap,attr"`
			SortMissingLast      string `xml:"sortMissingLast,attr"`
			OmitNorms            string `xml:"omitNorms,attr"`
			Indexed              string `xml:"indexed,attr"`
			Analyzer             []struct {
				Text      string `xml:",chardata"`
				Type      string `xml:"type,attr"`
				Tokenizer struct {
					Text    string `xml:",chardata"`
					Class   string `xml:"class,attr"`
					Pattern string `xml:"pattern,attr"`
					Group   string `xml:"group,attr"`
				} `xml:"tokenizer"`
				Filter []struct {
					Text                string `xml:",chardata"`
					Class               string `xml:"class,attr"`
					Pattern             string `xml:"pattern,attr"`
					Replacement         string `xml:"replacement,attr"`
					Replace             string `xml:"replace,attr"`
					GenerateWordParts   string `xml:"generateWordParts,attr"`
					GenerateNumberParts string `xml:"generateNumberParts,attr"`
					CatenateWords       string `xml:"catenateWords,attr"`
					CatenateNumbers     string `xml:"catenateNumbers,attr"`
					CatenateAll         string `xml:"catenateAll,attr"`
					SplitOnCaseChange   string `xml:"splitOnCaseChange,attr"`
					IgnoreCase          string `xml:"ignoreCase,attr"`
					Words               string `xml:"words,attr"`
					Protected           string `xml:"protected,attr"`
					Language            string `xml:"language,attr"`
					MaxShingleSize      string `xml:"maxShingleSize,attr"`
					OutputUnigrams      string `xml:"outputUnigrams,attr"`
					Min                 string `xml:"min,attr"`
					Max                 string `xml:"max,attr"`
				} `xml:"filter"`
				CharFilter struct {
					Text        string `xml:",chardata"`
					Class       string `xml:"class,attr"`
					Pattern     string `xml:"pattern,attr"`
					Replacement string `xml:"replacement,attr"`
				} `xml:"charFilter"`
			} `xml:"analyzer"`
		} `xml:"fieldType"`
	} `xml:"types"`
	Fields struct {
		Text  string `xml:",chardata"`
		Field []struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name,attr"`
			Type        string `xml:"type,attr"`
			Indexed     string `xml:"indexed,attr"`
			Stored      string `xml:"stored,attr"`
			MultiValued string `xml:"multiValued,attr"`
			TermVectors string `xml:"termVectors,attr"`
			Default     string `xml:"default,attr"`
		} `xml:"field"`
		DynamicField []struct {
			Text        string `xml:",chardata"`
			Name        string `xml:"name,attr"`
			Type        string `xml:"type,attr"`
			Indexed     string `xml:"indexed,attr"`
			Stored      string `xml:"stored,attr"`
			MultiValued string `xml:"multiValued,attr"`
		} `xml:"dynamicField"`
	} `xml:"fields"`
	UniqueKey          string `xml:"uniqueKey"`          // id
	DefaultSearchField string `xml:"defaultSearchField"` // allfields
	CopyField          []struct {
		Text   string `xml:",chardata"`
		Source string `xml:"source,attr"`
		Dest   string `xml:"dest,attr"`
	} `xml:"copyField"`
	SolrQueryParser struct {
		Text            string `xml:",chardata"`
		DefaultOperator string `xml:"defaultOperator,attr"`
	} `xml:"solrQueryParser"`
}

// GoName converts a string into a more idiomatic name. Might miss edge cases.
func GoName(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "-", "", -1)
	parts := strings.Split(s, "_")
	var camel []string
	for _, p := range parts {
		if stringSliceContains(UppercaseWords, strings.ToLower(p)) {
			camel = append(camel, strings.ToUpper(p))
		} else {
			camel = append(camel, strings.Title(p))
		}
	}
	return strings.Join(camel, "")
}

// RenderStringSlice renders a strings slice.
func RenderStringSlice(s []string) string {
	var buf bytes.Buffer
	io.WriteString(&buf, "[]string{\n")
	for _, v := range s {
		fmt.Fprintf(&buf, "%q,\n", v)
	}
	io.WriteString(&buf, "}")
	return buf.String()
}
