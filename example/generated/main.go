// Autogenerated by s2gen 0.1.1 on 2018-11-05T15:15:49+01:00 by tir on sol.
// Schema SHA1 was 405560333c6c79e90b6b07ff5c916c353741e92d.
// Do NOT modify, unless you know what you do.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/structs"
)

// VuFindBibliographicIndex generated from schema.
// This struct supports marshaling and unmarshaling of SOLR documents conforming to the
// schema that was used to generate this struct. Static fields are struct fields, dynamic
// fields can be modified by using v.MustField("dyn_*").Set(...) and so on.
type VuFindBibliographicIndex struct {
	Version              json.Number `json:"_version_"`
	ID                   string      `json:"id"`
	Fullrecord           string      `json:"fullrecord"`
	Itemdata             string      `json:"itemdata"`
	MARCError            []string    `json:"marc_error"`
	Allfields            []string    `json:"allfields"`
	AllfieldsUnstemmed   []string    `json:"allfields_unstemmed"`
	Fulltext             string      `json:"fulltext"`
	FulltextUnstemmed    string      `json:"fulltext_unstemmed"`
	Spelling             []string    `json:"spelling"`
	SpellingShingle      []string    `json:"spellingShingle"`
	AccessFacet          string      `json:"access_facet"`
	MegaCollection       []string    `json:"mega_collection"`
	RecordID             string      `json:"record_id"`
	SourceID             string      `json:"source_id"`
	Institution          []string    `json:"institution"`
	Collection           []string    `json:"collection"`
	Building             []string    `json:"building"`
	Language             []string    `json:"language"`
	Format               []string    `json:"format"`
	Signatur             []string    `json:"signatur"`
	Barcode              []string    `json:"barcode"`
	RSN                  []string    `json:"rsn"`
	Purchase             []string    `json:"purchase"`
	Timecode             []string    `json:"timecode"`
	MiscDech1            []string    `json:"misc_dech1"`
	FincClassFacet       []string    `json:"finc_class_facet"`
	ZDB                  string      `json:"zdb"`
	AuthorID             []string    `json:"author_id"`
	AuthorRef            []string    `json:"author_ref"`
	AuthorCorporateRef   []string    `json:"author_corporate_ref"`
	AuthorStr            string      `json:"authorStr"`
	AuthorFacet          []string    `json:"author_facet"`
	AuthorSort           string      `json:"author_sort"`
	Author               []string    `json:"author"`
	AuthorOrig           []string    `json:"author_orig"`
	AuthorRole           []string    `json:"author_role"`
	Author2              []string    `json:"author2"`
	Author2Orig          []string    `json:"author2_orig"`
	Author2Role          []string    `json:"author2_role"`
	AuthorCorporate      []string    `json:"author_corporate"`
	AuthorCorporateOrig  []string    `json:"author_corporate_orig"`
	AuthorCorporateRole  []string    `json:"author_corporate_role"`
	AuthorCorporate2     []string    `json:"author_corporate2"`
	AuthorCorporate2Orig []string    `json:"author_corporate2_orig"`
	AuthorCorporate2Role []string    `json:"author_corporate2_role"`
	AuthorAdditional     []string    `json:"author_additional"`
	Title                string      `json:"title"`
	TitlePart            string      `json:"title_part"`
	TitleSub             string      `json:"title_sub"`
	TitleShort           string      `json:"title_short"`
	TitleFull            string      `json:"title_full"`
	TitleFullUnstemmed   string      `json:"title_full_unstemmed"`
	TitleFullStr         string      `json:"title_fullStr"`
	TitleAuth            string      `json:"title_auth"`
	TitleAlt             []string    `json:"title_alt"`
	TitleOld             []string    `json:"title_old"`
	TitleNew             []string    `json:"title_new"`
	TitleUniform         string      `json:"title_uniform"`
	TitleSort            string      `json:"title_sort"`
	TitleOrig            string      `json:"title_orig"`
	Physical             []string    `json:"physical"`
	Publisher            []string    `json:"publisher"`
	PublisherStr         []string    `json:"publisherStr"`
	PublishDate          []string    `json:"publishDate"`
	PublishDateSort      string      `json:"publishDateSort"`
	PublishPlace         []string    `json:"publishPlace"`
	Imprint              string      `json:"imprint"`
	Edition              string      `json:"edition"`
	Description          string      `json:"description"`
	Contents             []string    `json:"contents"`
	URL                  []string    `json:"url"`
	Thumbnail            string      `json:"thumbnail"`
	LCCN                 string      `json:"lccn"`
	Ctrlnum              []string    `json:"ctrlnum"`
	URN                  []string    `json:"urn"`
	ISBN                 []string    `json:"isbn"`
	ISSN                 []string    `json:"issn"`
	ISMN                 []string    `json:"ismn"`
	OCLCNum              []string    `json:"oclc_num"`
	Callnumberfirst      string      `json:"callnumber-first"`
	Callnumbersubject    string      `json:"callnumber-subject"`
	Callnumberlabel      string      `json:"callnumber-label"`
	Callnumbersort       string      `json:"callnumber-sort"`
	Callnumberraw        []string    `json:"callnumber-raw"`
	Callnumbersearch     []string    `json:"callnumber-search"`
	Deweyhundreds        []string    `json:"dewey-hundreds"`
	Deweytens            []string    `json:"dewey-tens"`
	Deweyones            []string    `json:"dewey-ones"`
	Deweyfull            []string    `json:"dewey-full"`
	Deweysort            string      `json:"dewey-sort"`
	Deweyraw             []string    `json:"dewey-raw"`
	Deweysearch          []string    `json:"dewey-search"`
	DateSpan             []string    `json:"dateSpan"`
	Series               []string    `json:"series"`
	Series2              []string    `json:"series2"`
	SeriesOrig           []string    `json:"series_orig"`
	Topic                []string    `json:"topic"`
	TopicID              []string    `json:"topic_id"`
	TopicRef             []string    `json:"topic_ref"`
	TopicUnstemmed       []string    `json:"topic_unstemmed"`
	TopicFacet           []string    `json:"topic_facet"`
	TopicBrowse          []string    `json:"topic_browse"`
	AuthorBrowse         []string    `json:"author_browse"`
	Genre                []string    `json:"genre"`
	GenreFacet           []string    `json:"genre_facet"`
	Geographic           []string    `json:"geographic"`
	GeographicFacet      []string    `json:"geographic_facet"`
	GeogrCode            []string    `json:"geogr_code"`
	GeogrCodePerson      []string    `json:"geogr_code_person"`
	Era                  []string    `json:"era"`
	EraFacet             []string    `json:"era_facet"`
	Footnote             []string    `json:"footnote"`
	DissertationNote     []string    `json:"dissertation_note"`
	PerformerNote        []string    `json:"performer_note"`
	Illustrated          string      `json:"illustrated"`
	LongLat              string      `json:"long_lat"`
	MusicHeading         []string    `json:"music_heading"`
	MusicHeadingBrowse   []string    `json:"music_heading_browse"`
	FilmHeading          []string    `json:"film_heading"`
	RVKFacet             []string    `json:"rvk_facet"`
	RVKLabel             []string    `json:"rvk_label"`
	RVKPath              []string    `json:"rvk_path"`
	ContainerTitle       string      `json:"container_title"`
	ContainerVolume      string      `json:"container_volume"`
	ContainerIssue       string      `json:"container_issue"`
	ContainerStartPage   string      `json:"container_start_page"`
	ContainerReference   string      `json:"container_reference"`
	MultipartSet         string      `json:"multipart_set"`
	MultipartLink        []string    `json:"multipart_link"`
	MultipartPart        []string    `json:"multipart_part"`
	Hierarchytype        string      `json:"hierarchytype"`
	HierarchyTopID       []string    `json:"hierarchy_top_id"`
	HierarchyTopTitle    []string    `json:"hierarchy_top_title"`
	HierarchyParentID    []string    `json:"hierarchy_parent_id"`
	HierarchyParentTitle []string    `json:"hierarchy_parent_title"`
	HierarchySequence    []string    `json:"hierarchy_sequence"`
	IsHierarchyID        string      `json:"is_hierarchy_id"`
	IsHierarchyTitle     string      `json:"is_hierarchy_title"`
	TitleInHierarchy     []string    `json:"title_in_hierarchy"`
	HierarchyBrowse      []string    `json:"hierarchy_browse"`
	Recordtype           string      `json:"recordtype"`
	FirstIndexed         string      `json:"first_indexed"`
	LastIndexed          string      `json:"last_indexed"`

	fields []*dynamicField
}

// NewVuFindBibliographicIndex creates a new document conforming to a schema for the dynamic fields.
func NewVuFindBibliographicIndex() *VuFindBibliographicIndex {
	var v VuFindBibliographicIndex
	v.initDynamicFields()
	return &v
}

// initDynamicFields will initialize or reset the dynamic fields.
func (v *VuFindBibliographicIndex) initDynamicFields() {
	v.fields = []*dynamicField{
		&dynamicField{name: "callnumber_*", isMultiValued: true},
		&dynamicField{name: "barcode_*", isMultiValued: true},
		&dynamicField{name: "misc_*", isMultiValued: true},
		&dynamicField{name: "branch_*", isMultiValued: true},
		&dynamicField{name: "collcode_*", isMultiValued: true},
		&dynamicField{name: "format_*", isMultiValued: true},
		&dynamicField{name: "facet_*", isMultiValued: true},
		&dynamicField{name: "local_heading_facet_*", isMultiValued: true},
		&dynamicField{name: "local_heading_*", isMultiValued: true},
		&dynamicField{name: "local_class_*", isMultiValued: true},
		&dynamicField{name: "udk_raw_*", isMultiValued: true},
		&dynamicField{name: "udk_facet_*", isMultiValued: true},
		&dynamicField{name: "date_*", isMultiValued: false},
		&dynamicField{name: "*_date", isMultiValued: false},
		&dynamicField{name: "*_date_mv", isMultiValued: true},
		&dynamicField{name: "*_isn", isMultiValued: false},
		&dynamicField{name: "*_isn_mv", isMultiValued: true},
		&dynamicField{name: "*_str", isMultiValued: false},
		&dynamicField{name: "*_str_mv", isMultiValued: true},
		&dynamicField{name: "*_txt", isMultiValued: false},
		&dynamicField{name: "*_txt_mv", isMultiValued: true},
		&dynamicField{name: "*_txtF", isMultiValued: false},
		&dynamicField{name: "*_txtF_mv", isMultiValued: true},
		&dynamicField{name: "*_txtP", isMultiValued: false},
		&dynamicField{name: "*_txtP_mv", isMultiValued: true},
		&dynamicField{name: "*_random", isMultiValued: false},
		&dynamicField{name: "*_boolean", isMultiValued: false},
	}
}

// MarshalJSON serialized static and dynamic fields.
func (v *VuFindBibliographicIndex) MarshalJSON() ([]byte, error) {
	temp := structs.Map(v)
	for _, field := range v.fields {
		for key, value := range field.fmap {
			temp[key] = value
		}
	}
	return json.Marshal(temp)
}

// UnmarshalJSON unmarshals a document. It is not an error, if there are
// fields, which do not fit neither into static of dynamic definitions.  As
// unmarshaling can happen in larger structs, the dynamic fields are
// initialized as part of the process.
func (v *VuFindBibliographicIndex) UnmarshalJSON(pp []byte) error {
	v.initDynamicFields()
	temp := make(map[string]interface{})
	if err := json.Unmarshal(pp, &temp); err != nil {
		return err
	}
	st := structs.New(v)
	for _, field := range st.Fields() {
		name := field.Tag("json")
		if value, ok := temp[name]; ok {
			switch tv := value.(type) {
			case float64, string, []string:
				field.Set(tv)
			case []interface{}:
				field.Set(toStringSlice(tv))
			default:
				return fmt.Errorf("cannot set value for type %T", tv)
			}
		}
	}
	for _, field := range v.fields {
		for key, value := range temp {
			if err := field.isCompatible(key); err != nil {
				continue
			}
			if err := field.Set(key, value); err != nil {
				return err
			}
		}
	}
	return nil
}

// DynamicFieldnames returns the dynamic field wildcards.
func (v *VuFindBibliographicIndex) DynamicFieldnames() (names []string) {
	for _, f := range v.fields {
		names = append(names, f.name)
	}
	return
}

// DynamicFields returns a map of concrete fields for a given wildcard.
func (v *VuFindBibliographicIndex) DynamicFields(name string) (map[string][]string, error) {
	for _, f := range v.fields {
		if f.name == name {
			return f.fmap, nil
		}
	}
	return nil, fmt.Errorf("no such field: %s", name)
}

// Field returns the dynamic field with a given name. It is an error,
// if the field does not exist.
func (v *VuFindBibliographicIndex) Field(name string) (*dynamicField, error) {
	for _, field := range v.fields {
		if field.name == name {
			return field, nil
		}
	}
	return nil, fmt.Errorf("no such field: %s", name)
}

// MustField returns the field or nil.
func (v *VuFindBibliographicIndex) MustField(name string) *dynamicField {
	field, _ := v.Field(name)
	return field
}

// toStringSlice returns a slice of strings
func toStringSlice(is []interface{}) (values []string) {
	for _, v := range is {
		values = append(values, fmt.Sprintf("%s", v))
	}
	return
}

// dynamicField represents a more generic multi-valued field.
type dynamicField struct {
	name          string
	isMultiValued bool
	fmap          map[string][]string
}

// ensureMap lazily initializes the map.
func (f *dynamicField) ensureMap() {
	if f.fmap == nil {
		f.fmap = make(map[string][]string)
	}
}

// isCompatible returns nil, if the given key string can be used for this
// dynamic field.
func (f *dynamicField) isCompatible(key string) error {
	re, err := regexp.Compile("^" + strings.Replace(f.name, "*", ".*", -1) + "$")
	if err != nil {
		return err
	}
	if re.MatchString(key) {
		return nil
	}
	return fmt.Errorf("incompatible key")
}

// addStrings tries to add a list of strings to the given key k
func (f *dynamicField) addStrings(k string, ss []string) error {
	f.ensureMap()
	switch f.isMultiValued {
	case true:
		f.fmap[k] = append(f.fmap[k], ss...)
	case false:
		if _, ok := f.fmap[k]; ok {
			return fmt.Errorf("single-valued field %s exists", f.name)
		}
		if len(ss) > 1 {
			return fmt.Errorf("multiple values for single-valued field %s", f.name)
		}
		f.fmap[k] = ss
	}
	return nil
}

// setStrings tries to set a list of strings to the given key k
func (f *dynamicField) setStrings(k string, ss []string) error {
	f.ensureMap()
	switch f.isMultiValued {
	case true:
		f.fmap[k] = ss
	case false:
		if len(ss) > 1 {
			return fmt.Errorf("multiple values for single-valued field %s", f.name)
		}
		f.fmap[k] = ss
	}
	return nil
}

// Clear remove all values for a key.
func (f *dynamicField) Clear(k string) {
	f.ensureMap()
	if _, ok := f.fmap[k]; ok {
		delete(f.fmap, k)
	}
}

// Add value or values to a given field. It is an error if a single valued
// field would have more than one element.
func (f *dynamicField) Add(k string, v interface{}) (err error) {
	f.ensureMap()
	if err = f.isCompatible(k); err != nil {
		return
	}
	switch t := v.(type) {
	case string:
		err = f.addStrings(k, []string{t})
	case []string:
		err = f.addStrings(k, t)
	case []interface{}:
		err = f.addStrings(k, toStringSlice(t))
	default:
		err = fmt.Errorf("type %T for values not supported", v)
	}
	return
}

// Set sets the values for a given key. It accepts a single string or a string
// slice. If is an error, if a string slice is used for a single valued field.
func (f *dynamicField) Set(k string, v interface{}) (err error) {
	f.ensureMap()
	if err = f.isCompatible(k); err != nil {
		return
	}
	switch t := v.(type) {
	case string:
		f.fmap[k] = []string{t}
	case []string:
		err = f.setStrings(k, t)
	case []interface{}:
		err = f.setStrings(k, toStringSlice(t))
	default:
		err = fmt.Errorf("type %T for values not supported", v)
	}
	return
}

// Value returns a single value for a key and whether the key was actually present.
func (f *dynamicField) Value(key string) (value string, ok bool) {
	f.ensureMap()
	var vs []string
	vs, ok = f.fmap[key]
	if !ok || len(vs) == 0 {
		return
	}
	return vs[0], true
}

// Values returns multiple values for a key and whether the key was actually present.
func (f *dynamicField) Values(key string) (values []string, ok bool) {
	f.ensureMap()
	values, ok = f.fmap[key]
	if !ok {
		return
	}
	return values, true
}

// MustValue returns the value or an empty string.
func (f *dynamicField) MustValue(key string) string {
	value, _ := f.Value(key)
	return value
}

// MustValues returns slice of values of a nil slice.
func (f *dynamicField) MustValues(key string) []string {
	values, _ := f.Values(key)
	return values
}

func main() {
	dec := json.NewDecoder(os.Stdin)
	var doc VuFindBibliographicIndex

	if err := dec.Decode(&doc); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", doc)
}

